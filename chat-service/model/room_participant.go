package model

import "gorm.io/gorm"

type RoomParticipant struct {
	gorm.Model
	RoomID    uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	UserEmail string `gorm:"not null"`
	Room      Room   `gorm:"foreignKey:RoomID"` // Relasi ke Room
	// User   User `gorm:"foreignKey:UserID"` // Relasi ke User
}
