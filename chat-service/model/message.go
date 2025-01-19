package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	RoomID        uint       `json:"room_id"`
	SenderEmail   string     `json:"sender_email"`
	Content       string     `json:"content"`
	AttachmentURL *string    `json:"attachment_url"`
	ReplyTo       *uint      `json:"reply_to"`
	ReadAt        *time.Time `json:"read_at"`
	Room          Room       `gorm:"foreignKey:RoomID"` // Relasi ke Room
}
