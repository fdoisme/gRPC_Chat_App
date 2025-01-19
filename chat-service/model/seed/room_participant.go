package seed

import "project/chat-service/model"

func RoomParticipantSeed() []model.RoomParticipant {
	return []model.RoomParticipant{
		{RoomID: 1, UserEmail: "satu@mail.com"},
		{RoomID: 1, UserEmail: "dua@mail.com"},
		{RoomID: 2, UserEmail: "satu@mail.com"},
		{RoomID: 2, UserEmail: "tiga@mail.com"},
		{RoomID: 2, UserEmail: "empat@mail.com"},
	}
}
