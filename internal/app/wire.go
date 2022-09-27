//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"maxblog-me-template/internal/api"
	"maxblog-me-template/internal/conf"
	"maxblog-me-template/internal/core"
	"maxblog-me-template/internal/middleware/interceptor"
	"maxblog-me-template/src/handler"
	"maxblog-me-template/src/service"
)

func InitInjector() (*Injector, error) {
	wire.Build(
		conf.InitGinEngine,
		api.APISet,
		interceptor.InterceptorSet,
		core.ResponseSet,
		core.LoggerSet,
		handler.HandlerSet,
		service.ServiceSet,
		InjectorSet,
	)
	return new(Injector), nil
}
