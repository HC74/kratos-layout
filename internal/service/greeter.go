package service

import (
	"context"
	"github.com/HC74/kratosx"
	v1 "kratos-layout/api/v1"
)

func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReplyTest, error) {
	return s.logic.SayHello(kratosx.MustContext(ctx), in)
}
