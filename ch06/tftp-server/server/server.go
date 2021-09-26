package server

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"tftp-server/tftp"
	"time"
)

var ErrRetryable = errors.New("something went wrong but can be retried")

func New(payload []byte, retries uint8, timeout time.Duration) (*Server, error) {
	if payload == nil {
		return nil, errors.New("payload is nil")
	}

	return &Server{
		payload: payload,
		retries: retries,
		timeout: timeout,
	}, nil
}

type Server struct {
	payload []byte
	retries uint8
	timeout time.Duration
}

func (s *Server) ListenAndServe(addr string) error {
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return fmt.Errorf("couldn't start server: %w", err)
	}
	defer conn.Close()

	buf := make([]byte, tftp.DatagramSize)
	for {
		_, clientAddr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Println("couldn't read from client connection: " + err.Error())
			continue
		}

		fmt.Println("inc request")
		rrq, err := tftp.NewReadRequestFromBinary(buf)
		if err != nil {
			log.Println("couldn't unmarshal read request: " + err.Error())
		}

		fmt.Printf("handling request for client address %s\n", clientAddr.String())
		s.handleRequest(conn, clientAddr, rrq)
	}
}

func (s *Server) handleRequest(conn net.PacketConn, clientAddr net.Addr, rrq *tftp.ReadRequest) {
	dataPacket := tftp.NewData(0, bytes.NewBuffer(s.payload))
	currBlockSize := tftp.DatagramSize
	for currBlockSize == tftp.DatagramSize {
		data, err := dataPacket.MarshalBinary()
		if err != nil {
			log.Println("couldn't marshal data packet: " + err.Error())
			return
		}

		err = withRetries(s.retries, sendData(s.timeout, conn, clientAddr, data, dataPacket.Block))
		if err != nil {
			log.Println("unexpected error: " + err.Error())
			return
		}

		fmt.Printf("block #%d sent successfully!\n", dataPacket.Block)
		currBlockSize = len(data)
	}

	fmt.Println("file successfully transferred!")
}

func sendData(timeout time.Duration, conn net.PacketConn, clientAddr net.Addr, data []byte, expectedBlock uint16) func() error {
	return func() error {
		fmt.Printf("attempting to send data block #%d...\n", expectedBlock)

		_, err := conn.WriteTo(data, clientAddr)
		if err != nil {
			return err
		}

		err = conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			return err
		}

		resp := make([]byte, tftp.DatagramSize)
		_, _, err = conn.ReadFrom(resp)
		if err != nil {
			return err
		}

		if ack, err := tftp.NewAckFromBinary(resp); err == nil {
			if ack.Block == expectedBlock {
				return nil
			}

			return ErrRetryable
		} else if eResp, err := tftp.NewErrFromBinary(resp); err == nil {
			log.Printf("[%s] received error: [code: %d, msg: %v]", clientAddr.String(), eResp.Code, eResp.Message)
			return errors.New(eResp.Message)
		} else {
			log.Printf("[%s] bad packet", clientAddr.String())
			return ErrRetryable
		}
	}
}

func withRetries(max uint8, f func() error) error {
	for i := int(max); i > 0; i-- {
		err := f()
		if nErr, ok := err.(net.Error); ok && nErr.Temporary() {
			fmt.Printf("received transient network error: %v", err)
			continue
		} else if err == ErrRetryable {
			continue
		}

		return err
	}

	return errors.New("max retries exceeded")
}
