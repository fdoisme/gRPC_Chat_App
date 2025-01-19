package seed

import "project/chat-service/model"

func RoomSeed() []model.Room {
	return []model.Room{
		{Name: "General Room"},
		{Name: "Support Room"},
		{Name: "Private Room 1"},
		{Name: "Private Room 2"},
	}
}
