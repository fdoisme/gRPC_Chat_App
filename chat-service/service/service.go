package service

import (
	"project/chat-service/repository"

	"go.uber.org/zap"
)

type Service struct {
	ChatService ChatService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		ChatService: NewChatService(repo),
	}
}
