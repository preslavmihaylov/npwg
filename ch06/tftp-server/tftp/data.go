package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

func NewData(block uint16, payload io.Reader) *Data {
	return &Data{
		Block:   block,
		Payload: payload,
	}
}

func NewDataFromBinary(data []byte) (*Data, error) {
	d := &Data{}
	err := d.UnmarshalBinary(data)
	return d, err
}

type Data struct {
	Block   uint16
	Payload io.Reader
}

func (d *Data) MarshalBinary() (data []byte, err error) {
	opCode := OpData

	d.Block++
	block := d.Block

	buffer := new(bytes.Buffer)
	buffer.Grow(DatagramSize)

	err = allMust(
		binary.Write(buffer, binary.BigEndian, opCode),
		binary.Write(buffer, binary.BigEndian, block),
		copyN(buffer, d.Payload, BlockSize),
	)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (d *Data) UnmarshalBinary(data []byte) error {
	var opCode OpCode
	var block uint16
	dataBlock := bytes.NewBuffer([]byte{})

	buffer := bytes.NewBuffer(data)
	err := allMust(
		binary.Read(buffer, binary.BigEndian, &opCode),
		binary.Read(buffer, binary.BigEndian, &block),
		copyN(dataBlock, buffer, BlockSize),
	)
	if err != nil {
		return err
	} else if opCode != OpData {
		return errors.New("invalid data datagram")
	}

	d.Block, d.Payload = block, dataBlock
	return nil
}
