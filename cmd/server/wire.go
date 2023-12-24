//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/HC74/kratos-layout/internal/biz"
	"github.com/HC74/kratos-layout/internal/conf"
	"github.com/HC74/kratos-layout/internal/data"
	"github.com/HC74/kratos-layout/internal/server"
	"github.com/HC74/kratos-layout/internal/service"

	"github.com/HC74/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
