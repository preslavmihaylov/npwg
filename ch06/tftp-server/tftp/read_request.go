package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func NewReadRequest(filename, mode string) *ReadRequest {
	return &ReadRequest{
		Filename: filename,
		Mode:     mode,
	}
}

func NewReadRequestFromBinary(data []byte) (*ReadRequest, error) {
	rr := &ReadRequest{}
	err := rr.UnmarshalBinary(data)
	return rr, err
}

type ReadRequest struct {
	Filename string
	Mode     string
}

func (rr *ReadRequest) MarshalBinary() ([]byte, error) {
	opCode := OpRRQ
	mode := rr.Mode
	if mode == "" {
		mode = "octet"
	}

	buffer := bytes.NewBuffer([]byte{})

	err := allMust(
		binary.Write(buffer, binary.BigEndian, opCode),
		binary.Write(buffer, binary.BigEndian, []byte(rr.Filename)),
		binary.Write(buffer, binary.BigEndian, []byte{0}),
		binary.Write(buffer, binary.BigEndian, []byte(mode)),
		binary.Write(buffer, binary.BigEndian, []byte{0}),
	)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (rr *ReadRequest) UnmarshalBinary(data []byte) error {
	var opCode OpCode
	var filename string
	var mode string

	buffer := bytes.NewBuffer(data)
	err := allMust(
		binary.Read(buffer, binary.BigEndian, &opCode),
		readString(buffer, &filename, "\x00"),
		readString(buffer, &mode, "\x00"),
	)
	if err != nil {
		return err
	} else if !isValidReadReq(opCode, filename, mode) {
		return errors.New("invalid read request")
	}

	rr.Filename, rr.Mode = filename, mode
	return nil
}

func isValidReadReq(opCode OpCode, filename, mode string) bool {
	return opCode == OpRRQ && mode == "octet"
}
