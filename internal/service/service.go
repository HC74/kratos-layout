package service

import (
	v1 "kratos-layout/api/v1"
	super "kratos-layout/config"
	"kratos-layout/internal/logic"
)

// Service is a greeter service.
type Service struct {
	v1.UnimplementedServiceServer
	logic *logic.Logic
}

func New(conf *super.Config) *Service {
	return &Service{
		logic: logic.NewLogic(conf),
	}
}
