package lib

import (
	"encoding/binary"
	"io"
)

func AsBinary(data []byte) *Binary {
	b := Binary(data)
	return &b
}

type Binary []byte

func (bin *Binary) String() string {
	return string(*bin)
}

func (bin *Binary) ReadFrom(r io.Reader) (int64, error) {
	var theType uint8
	err := binary.Read(r, binary.BigEndian, &theType)
	if err != nil {
		return 0, err
	}

	var theSize uint32
	err = binary.Read(r, binary.BigEndian, &theSize)
	if err != nil {
		return 0, err
	}

	*bin = make([]byte, theSize)
	n, err := r.Read(*bin)
	if err != nil {
		return 0, err
	}

	return int64(5 + n), nil
}

func (bin *Binary) WriteTo(w io.Writer) (int64, error) {
	var n int64 = 1
	err := binary.Write(w, binary.BigEndian, BinaryType)
	if err != nil {
		return 0, err
	}

	n += 4
	err = binary.Write(w, binary.BigEndian, uint32(len(*bin)))
	if err != nil {
		return 0, err
	}

	written, err := w.Write(*bin)
	if err != nil {
		return 0, err
	}

	return n + int64(written), nil
}

func (bin *Binary) Bytes() []byte {
	return *bin
}
