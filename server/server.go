package server

import (
	"log"
	"net"

	"github.com/object88/isomorphicTest"
	"github.com/object88/isomorphicTest/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9876"
)

// InitializeService starts the service
func InitializeService() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	s := grpc.NewServer()
	proto.RegisterGeneratorServer(s, &isomorphicTest.Generator{S: s})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	return s.Serve(lis)
}
