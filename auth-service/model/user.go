package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Email      string `gorm:"not null"`
	VerifiedAt *time.Time
	gorm.Model
}

func UserEmailUniqueIndex() string {
	return "CREATE UNIQUE INDEX unique_verified_email ON users (email) WHERE verified_at IS NOT NULL AND deleted_at IS NULL"
}
