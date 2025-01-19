package infra

import (
	"go.uber.org/zap"
	"project/api-gateway/config"
	"project/api-gateway/database"
	"project/api-gateway/handler"
	"project/api-gateway/log"
	"project/api-gateway/middleware"
	"project/api-gateway/service"
)

type ServiceContext struct {
	Cacher     database.Cacher
	Cfg        config.Config
	Ctl        handler.Handler
	Log        *zap.Logger
	Middleware middleware.Middleware
	Svc        *service.Service
}

func NewServiceContext() (*ServiceContext, error) {
	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	appConfig, err := config.LoadConfig()
	if err != nil {
		return handlerError(err)
	}

	// instance logger
	logger, err := log.InitZapLogger(appConfig)
	if err != nil {
		return handlerError(err)
	}

	rdb := database.NewCacher(appConfig, 60*60)

	// instance service
	services := service.NewService(appConfig, logger)

	// instance controller
	Ctl := handler.NewHandler(services, logger, rdb)

	mw := middleware.NewMiddleware(appConfig.MicroserviceConfig, rdb)

	//return &ServiceContext{Cacher: rdb, Cfg: appConfig, Svc: &services, Log: logger}, nil
	return &ServiceContext{Cfg: appConfig, Ctl: *Ctl, Middleware: mw, Log: logger, Svc: &services}, nil
}
