package middleware

import (
	"project/api-gateway/config"
	"project/api-gateway/database"
)

type Middleware struct {
	cacher         database.Cacher
	authServiceUrl string
}

func NewMiddleware(serviceConfig config.MicroserviceConfig, cacher database.Cacher) Middleware {
	return Middleware{authServiceUrl: serviceConfig.Auth, cacher: cacher}
}
