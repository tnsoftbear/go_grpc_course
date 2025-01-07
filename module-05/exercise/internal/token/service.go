package token

import (
	"context"
	"fmt"

	"github.com/cshep4/grpc-course/module5-exercise/proto"
)

type Service struct {
	proto.UnimplementedTokenServiceServer
}

func (s Service) Validate(ctx context.Context, _ *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	fmt.Println("Start validation")
	return &proto.ValidateResponse{}, nil
}
