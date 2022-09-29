// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"maxblog-me-template/internal/api"
	"maxblog-me-template/internal/conf"
	"maxblog-me-template/internal/core"
	"maxblog-me-template/internal/middleware/interceptor"
	"maxblog-me-template/src/handler"
	"maxblog-me-template/src/service"
)

// Injectors from wire.go:

func InitInjector() (*Injector, error) {
	bData := &service.BData{}
	logger := &core.Logger{}
	response := &core.Response{
		ILogger: logger,
	}
	hData := &handler.HData{
		BData: bData,
		IRes:  response,
	}
	apiHandler := &api.Handler{
		HandlerData: hData,
	}
	auth := &interceptor.Auth{
		ILogger: logger,
	}
	interceptorInterceptor := &interceptor.Interceptor{
		InterceptorAuth: auth,
	}
	engine := conf.InitGinEngine(apiHandler, interceptorInterceptor)
	injector := &Injector{
		Engine:      engine,
		Handler:     apiHandler,
		Interceptor: interceptorInterceptor,
	}
	return injector, nil
}
