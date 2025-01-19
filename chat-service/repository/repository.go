package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	ChatRepo ChatRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		ChatRepo: NewChatRepository(db, log),
	}
}
