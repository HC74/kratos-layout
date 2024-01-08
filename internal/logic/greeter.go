package logic

import (
	"github.com/HC74/kratosx"
	v1 "kratos-layout/api/v1"
)

// SayHello implements SayHello
func (l *Logic) SayHello(ctx kratosx.Context, in *v1.HelloRequest) (*v1.HelloReplyTest, error) {
	// TODO Logic业务逻辑
	return &v1.HelloReplyTest{
		Message: l.conf.Msg,
	}, nil
}
