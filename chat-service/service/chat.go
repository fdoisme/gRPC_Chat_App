package service

import (
	"project/chat-service/model"
	"project/chat-service/repository"
)

type ChatService interface {
	GetUserDetails(userID uint) (*model.User, error)
	CreateRoom(room *model.Room) error
	CreateRoomParticipant(roomParticipant *model.RoomParticipant) error
	SaveMessage(message *model.Message) error
	GetRoomParticipants(roomID uint) ([]model.RoomParticipant, error)
	GetRoomMessages(roomID uint, limit int, page int) (*model.Pagination, error)
	GetRoomDetails(roomID uint) (*model.Room, error)
}

type chatService struct {
	repo repository.Repository
}

func NewChatService(repo repository.Repository) ChatService {
	return &chatService{repo: repo}
}

func (s *chatService) GetUserDetails(userID uint) (*model.User, error) {
	return s.repo.ChatRepo.GetUserDetails(userID)
}

func (s *chatService) CreateRoom(room *model.Room) error {
	return s.repo.ChatRepo.CreateRoom(room)
}

func (s *chatService) CreateRoomParticipant(roomParticipant *model.RoomParticipant) error {
	return s.repo.ChatRepo.CreateRoomParticipant(roomParticipant)
}

func (s *chatService) SaveMessage(message *model.Message) error {
	return s.repo.ChatRepo.SaveMessage(message)
}

func (s *chatService) GetRoomParticipants(roomID uint) ([]model.RoomParticipant, error) {
	return s.repo.ChatRepo.GetRoomParticipants(roomID)
}

func (s *chatService) GetRoomMessages(roomID uint, limit int, page int) (*model.Pagination, error) {
	offset := (page - 1) * limit
	return s.repo.ChatRepo.GetRoomMessages(roomID, limit, offset)
}

func (s *chatService) GetRoomDetails(roomID uint) (*model.Room, error) {
	return s.repo.ChatRepo.GetRoomByID(roomID)
}
