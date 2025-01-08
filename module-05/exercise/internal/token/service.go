package token

import (
	"context"
	"errors"
	"github.com/cshep4/grpc-course/module5-exercise/proto"
)

const ctxKeyClaims = "claims"

type Service struct {
	proto.UnimplementedTokenServiceServer
}

func (s Service) Validate(ctx context.Context, _ *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	claims, ok := ctx.Value(ctxKeyClaims).(map[string]string)
	if !ok {
		return nil, errors.New("claims missed in context")
	}

	return &proto.ValidateResponse{
		Claims: claims,
	}, nil
}
