package repository

import (
	"project/api-gateway/config"
	"project/api-gateway/database"

	"go.uber.org/zap"
)

type Repository struct {
}

func NewRepository(cacher database.Cacher, config config.Config, log *zap.Logger) Repository {
	return Repository{}
}
