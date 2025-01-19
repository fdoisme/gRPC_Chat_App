package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name         string            `json:"name" gorm:"default:pv"`
	Participants []RoomParticipant `json:"participants" gorm:"foreignKey:RoomID"`
	Messages     []Message         `json:"messages" gorm:"foreignKey:RoomID"`
}
