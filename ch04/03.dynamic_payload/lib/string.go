package lib

import (
	"encoding/binary"
	"errors"
	"io"
)

func AsString(data []byte) *String {
	s := String(data)
	return &s
}

type String []byte

func (s *String) String() string {
	return string(*s)
}

func (s *String) ReadFrom(r io.Reader) (int64, error) {
	var theType uint8
	err := binary.Read(r, binary.BigEndian, &theType)
	if err != nil {
		return 0, err
	} else if theType != StringType {
		return 0, errors.New("invalid type")
	}

	var theSize uint32
	err = binary.Read(r, binary.BigEndian, &theSize)
	if err != nil {
		return 0, err
	} else if theSize > MaxPayloadSize {
		return 0, errors.New("exceeded maximum payload size")
	}

	*s = make([]byte, theSize)
	n, err := r.Read(*s)
	if err != nil {
		return 0, err
	}

	return int64(5 + n), nil
}

func (s *String) WriteTo(w io.Writer) (int64, error) {
	var n int64 = 1
	err := binary.Write(w, binary.BigEndian, StringType)
	if err != nil {
		return 0, err
	}

	n += 4
	err = binary.Write(w, binary.BigEndian, uint32(len(*s)))
	if err != nil {
		return 0, err
	}

	written, err := w.Write(*s)
	if err != nil {
		return 0, err
	}

	return n + int64(written), nil
}

func (s *String) Bytes() []byte {
	return *s
}
