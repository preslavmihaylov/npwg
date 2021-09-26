package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func NewAck(block uint16) *Ack {
	return &Ack{
		Block: block,
	}
}

func NewAckFromBinary(data []byte) (*Ack, error) {
	ack := &Ack{}
	err := ack.UnmarshalBinary(data)
	return ack, err
}

type Ack struct {
	Block uint16
}

func (ack *Ack) MarshalBinary() ([]byte, error) {
	opCode := OpAck

	buffer := new(bytes.Buffer)
	buffer.Grow(4)

	err := allMust(
		binary.Write(buffer, binary.BigEndian, opCode),
		binary.Write(buffer, binary.BigEndian, ack.Block),
	)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (ack *Ack) UnmarshalBinary(data []byte) error {
	var opCode OpCode
	var block uint16

	buffer := bytes.NewBuffer(data)
	err := allMust(
		binary.Read(buffer, binary.BigEndian, &opCode),
		binary.Read(buffer, binary.BigEndian, &block),
	)
	if err != nil {
		return err
	} else if opCode != OpAck {
		return errors.New("invalid ack datagram")
	}

	ack.Block = block
	return nil
}
