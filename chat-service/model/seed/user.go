package seed

import "project/chat-service/model"

func UserSeed() []model.User {
	return []model.User{
		{Username: "user1", Email: "user1@example.com", Password: "password123"},
		{Username: "user2", Email: "user2@example.com", Password: "password123"},
		{Username: "user3", Email: "user3@example.com", Password: "password123"},
		{Username: "user4", Email: "user4@example.com", Password: "password123"},
	}
}
