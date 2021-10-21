package server

import (
	"log"
	"net"

	pbgen "housework/idl"

	"google.golang.org/grpc"
)

func Serve(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rosie := &Rosie{}
	server := grpc.NewServer()
	pbgen.RegisterRobotMaidServer(server, rosie)

	log.Printf("server listening at %v", lis.Addr())
	return server.Serve(lis)
}
