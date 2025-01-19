package service

import (
	"project/user-service/config"
	"project/user-service/repository"

	"go.uber.org/zap"
)

type Service struct {
	User *UserService
}

func NewService(repo repository.Repository, appConfig config.Config, log *zap.Logger) Service {
	return Service{
		User: NewUserService(repo, log),
	}
}
