package model

import "time"

type Message struct {
	Id            uint      `json:"id,-"`
	RoomId        uint      `json:"roomId,-"`
	Sender        string    `json:"sender,omitempty"`
	Content       string    `json:"content,omitempty"`
	AttachmentUrl string    `json:"attachmentUrl,omitempty"`
	ReplyTo       int       `json:"replyTo,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

type RoomParticipant struct {
	RoomId   uint
	RoomName string
	Users    []string
}

type Messages struct {
	RoomId   uint64
	RoomName string
	Messages []Message
}
type Participant struct {
	Email string `json:"email"`
}
