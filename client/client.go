package client

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/object88/isomorphicTest/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	address = "localhost:9876"
)

// Client is a wrapper for the gRPC client that handles connections
type Client struct {
	conn *grpc.ClientConn
	c    proto.GeneratorClient
}

// NewClient returns a new client
func NewClient() (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to make dial.\n")
		return nil, err
	}

	c := proto.NewGeneratorClient(conn)

	return &Client{conn, c}, nil
}

// DestroyClient tears down the gRPC connection
func (c *Client) DestroyClient() {
	c.conn.Close()
}

// RequestNewService will spin up a new server-oriented process and
// return a new client with a connection to it
func (c *Client) RequestNewService() (*Client, error) {
	c.DestroyClient()

	cmd := exec.Command(os.Args[0], "serve")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to spin up binary!\n")
		return nil, err
	}

	time.Sleep(500 * time.Millisecond)

	return NewClient()
}

// RequestShutdown sends a gRPC call to request the service process to shut down.
func (c *Client) RequestShutdown() error {
	_, err := c.c.Shutdown(context.Background(), &proto.ShutdownRequest{})
	if err == nil {
		return nil
	}

	s, ok := status.FromError(err)
	if !ok {
		return err
	}

	code := s.Code()
	if code == codes.Unavailable {
		return nil
	}

	return err
}

// GenerateUUID will create a new UUID
func (c *Client) GenerateUUID() (string, bool) {
	r, err := c.c.GenerateUUID(context.Background(), &proto.UUIDRequest{})
	if err == nil {
		return r.Uuid, true
	}

	s, ok := status.FromError(err)
	if !ok {
		return "", true
	}

	code := s.Code()
	return "", code != codes.Unavailable
}
