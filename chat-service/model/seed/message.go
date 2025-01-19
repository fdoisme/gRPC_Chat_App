package seed

import (
	"project/chat-service/helper"
	"project/chat-service/model"
	"time"
)

func MessageSeed() []model.Message {
	return []model.Message{
		newMessage(1, "satu@mail.com", "Welcome to the General Room!", helper.DateTime("2024-12-28 10:10:01")),
		newMessage(1, "satu@mail.com", "Lorem Ipsum", helper.DateTime("2024-12-28 10:10:01"), helper.DateTime("2024-12-28 14:15:11")),
		newMessage(1, "dua@mail.com", "Hello, how can I assist you today?", helper.DateTime("2024-12-28 10:10:01")),
		newMessage(2, "satu@mail.com", "I need help with my account.", helper.DateTime("2024-12-28 10:10:01")),
		newMessage(2, "tiga@mail.com", "Hi, this is a message.", helper.DateTime("2024-12-28 10:10:01")),
		newMessage(2, "empat@mail.com", "Room 2, message from User empat@mail.com.!", helper.DateTime("2024-12-28 10:10:01")),
	}
}

func newMessage(roomID uint, senderEmail string, content string, createdAt time.Time, readAt ...time.Time) model.Message {
	message := model.Message{RoomID: roomID, SenderEmail: senderEmail, Content: content}
	message.CreatedAt = createdAt
	if readAt != nil {
		message.ReadAt = helper.Ptr(readAt[0])
	}
	return message
}
