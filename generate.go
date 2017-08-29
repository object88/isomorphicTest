package isomorphicTest

import (
	"context"

	"github.com/google/uuid"
	"github.com/object88/isomorphicTest/proto"
)

type Generator struct{}

// GenerateUUID returns a string UUID
func (g *Generator) GenerateUUID(ctx context.Context, _ *proto.UUIDRequest) (*proto.UUIDReply, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	resp := &proto.UUIDReply{
		Uuid: u.String(),
	}

	return resp, nil
}
