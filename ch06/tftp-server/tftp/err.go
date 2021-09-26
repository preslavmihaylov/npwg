package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func NewErr(code ErrCode, msg string) *Err {
	return &Err{
		Code:    code,
		Message: msg,
	}
}

func NewErrFromBinary(data []byte) (*Err, error) {
	e := &Err{}
	err := e.UnmarshalBinary(data)
	return e, err
}

type Err struct {
	Code    ErrCode
	Message string
}

func (e *Err) MarshalBinary() ([]byte, error) {
	opCode := OpErr

	buffer := bytes.NewBuffer([]byte{})
	err := allMust(
		binary.Write(buffer, binary.BigEndian, opCode),
		binary.Write(buffer, binary.BigEndian, e.Code),
		binary.Write(buffer, binary.BigEndian, []byte(e.Message)),
		binary.Write(buffer, binary.BigEndian, []byte{0}),
	)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (e *Err) UnmarshalBinary(data []byte) error {
	var opCode OpCode
	var errCode ErrCode
	var msg string

	buffer := bytes.NewBuffer(data)
	err := allMust(
		binary.Read(buffer, binary.BigEndian, &opCode),
		binary.Read(buffer, binary.BigEndian, &errCode),
		readString(buffer, &msg, "\x00"),
	)
	if err != nil {
		return err
	} else if opCode != OpErr || errCode >= ErrCodeCnt {
		return errors.New("invalid read request")
	}

	e.Code, e.Message = errCode, msg
	return nil
}
