// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"juice/app/user/internal/biz"
	"juice/app/user/internal/conf"
	"juice/app/user/internal/data"
	"juice/app/user/internal/server"
	"juice/app/user/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewDB(confData)
	redisClient := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, client, redisClient)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	userBasicRepo := data.NewUserBasicRepo(dataData, logger)
	userBasic := biz.NewUserBasic(userBasicRepo,logger)
	userBasicService := service.NewUserBasicService(userBasic)
	grpcServer := server.NewGRPCServer(confServer, greeterService, userBasicService, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, userBasicService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
