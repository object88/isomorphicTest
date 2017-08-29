package isomorphicTest

import (
	"context"

	"github.com/google/uuid"
	"github.com/object88/isomorphicTest/proto"
	"google.golang.org/grpc"
)

// Generator keeps the gRPC Server reference
type Generator struct {
	S *grpc.Server
}

// GenerateUUID returns a string UUID
func (g *Generator) GenerateUUID(ctx context.Context, _ *proto.EmptyRequest) (*proto.UUIDReply, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	resp := &proto.UUIDReply{
		Uuid: u.String(),
	}

	return resp, nil
}

// Shutdown stops the service process
func (g *Generator) Shutdown(_ context.Context, _ *proto.EmptyRequest) (*proto.EmptyReply, error) {
	g.S.GracefulStop()
	return &proto.EmptyReply{}, nil
}

// Startup is a no-op to start the service
func (g *Generator) Startup(_ context.Context, _ *proto.EmptyRequest) (*proto.EmptyReply, error) {
	return &proto.EmptyReply{}, nil
}
