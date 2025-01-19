package handler

import (
	"go.uber.org/zap"
	"project/chat-service/service"
)

type Handler struct {
	ChatHandler ChatHandler
}

func NewHandler(service service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		ChatHandler: *NewChatHandler(service, logger),
	}
}
