package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"example/lib"
	"fmt"
	"io"
	"net"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	payload, err := decodeFrom(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println(payload.String())
}

func decodeFrom(r io.Reader) (lib.Payload, error) {
	var theType uint8
	err := binary.Read(r, binary.BigEndian, &theType)
	if err != nil {
		return nil, err
	}

	switch theType {
	case lib.BinaryType:
		bin := &lib.Binary{}
		_, err = bin.ReadFrom(io.MultiReader(bytes.NewBuffer([]byte{theType}), r))
		if err != nil {
			return nil, err
		}

		return bin, nil
	case lib.StringType:
		s := &lib.String{}
		_, err = s.ReadFrom(io.MultiReader(bytes.NewBuffer([]byte{theType}), r))
		if err != nil {
			return nil, err
		}

		return s, nil
	default:
		return nil, errors.New("unknown type: " + strconv.Itoa(int(theType)))

	}

}
