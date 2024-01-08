package main

import (
	"github.com/HC74/kratosx"
	"github.com/HC74/kratosx/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"
	v1 "kratos-layout/api/v1"
	super "kratos-layout/config"
	"kratos-layout/internal/service"
)

func main() {
	app := kratosx.New(
		kratosx.RegisterServer(RegisterServer),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	conf := &super.Config{}
	// config/config.yaml -> customize
	if err := c.Value("customize").Scan(conf); err != nil {
		panic("business config format error:" + err.Error())
	}
	srv := service.New(conf)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)
}
