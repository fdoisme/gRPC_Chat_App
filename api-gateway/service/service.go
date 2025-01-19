package service

import (
	"go.uber.org/zap"
	"project/api-gateway/config"
)

type Service struct {
	Auth  AuthService
	Chat  ChatService
	Email EmailService
	User  UserService
}

func NewService(appConfig config.Config, log *zap.Logger) Service {
	return Service{
		Auth:  NewAuthService(appConfig.MicroserviceConfig.Auth, log),
		Chat:  NewChatService(appConfig.MicroserviceConfig.Chat, log),
		Email: NewEmailService(appConfig.Email, log),
		User:  NewUserService(appConfig.MicroserviceConfig.User, log),
	}
}
