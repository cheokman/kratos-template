//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-template/app/helloworld/service/internal/biz"
	"kratos-template/app/helloworld/service/internal/conf"
	"kratos-template/app/helloworld/service/internal/data"
	"kratos-template/app/helloworld/service/internal/server"
	"kratos-template/app/helloworld/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
