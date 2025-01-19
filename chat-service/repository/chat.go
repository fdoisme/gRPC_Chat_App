package repository

import (
	"project/chat-service/config"
	"project/chat-service/database"
	"project/chat-service/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ChatRepository interface {
	GetUserDetails(userID uint) (*model.User, error)
	CreateRoom(room *model.Room) error
	CreateRoomParticipant(roomParticipant *model.RoomParticipant) error
	SaveMessage(message *model.Message) error
	GetRoomParticipants(roomID uint) ([]model.RoomParticipant, error)
	GetRoomMessages(roomID uint, limit int, offset int) (*model.Pagination, error)
	GetRoomByID(roomID uint) (*model.Room, error)
}

type chatRepository struct {
	DB     *gorm.DB
	Cacher database.Cacher
	Config config.Config
	Log    *zap.Logger
}

func NewChatRepository(db *gorm.DB, log *zap.Logger) ChatRepository {
	return &chatRepository{
		DB:  db,
		Log: log,
	}
}

func (r *chatRepository) GetUserDetails(userID uint) (*model.User, error) {
	var user model.User
	err := r.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *chatRepository) CreateRoom(room *model.Room) error {
	return r.DB.Create(room).Error
}

func (r *chatRepository) CreateRoomParticipant(roomParticipant *model.RoomParticipant) error {
	return r.DB.Create(roomParticipant).Error
}

func (r *chatRepository) SaveMessage(message *model.Message) error {
	return r.DB.Create(message).Error
}

func (r *chatRepository) GetRoomParticipants(roomID uint) ([]model.RoomParticipant, error) {
	var participants []model.RoomParticipant
	if err := r.DB.Where("room_id = ?", roomID).Find(&participants).Error; err != nil {
		return nil, err
	}
	return participants, nil
}

func (r *chatRepository) GetRoomMessages(roomID uint, limit int, offset int) (*model.Pagination, error) {
	var messages []model.Message
	var totalItems int64

	// Query to get the total count of messages in the room
	if err := r.DB.Model(&model.Message{}).Where("room_id = ?", roomID).Count(&totalItems).Error; err != nil {
		return nil, err
	}

	// Query to get the paginated messages
	if err := r.DB.Where("room_id = ?", roomID).Order("created_at desc").Limit(limit).Offset(offset).Find(&messages).Error; err != nil {
		return nil, err
	}

	// Calculate total pages
	totalPages := int(totalItems) / limit
	if totalItems%int64(limit) != 0 {
		totalPages++
	}

	// Return the pagination structure
	pagination := &model.Pagination{
		Page:       offset/limit + 1, // Current page (1-based)
		Limit:      limit,
		TotalItems: int(totalItems),
		TotalPages: totalPages,
		Messages:   messages,
	}

	return pagination, nil
}

// func (r *chatRepository) GetRoomMessages(roomID uint, limit int, offset int) ([]model.Message, error) {
// 	var messages []model.Message
// 	if err := r.DB.Where("room_id = ?", roomID).Order("created_at desc").Limit(limit).Offset(offset).Find(&messages).Error; err != nil {
// 		return nil, err
// 	}
// 	return messages, nil
// }

func (r *chatRepository) GetRoomByID(roomID uint) (*model.Room, error) {
	var room model.Room
	if err := r.DB.First(&room, roomID).Error; err != nil {
		return nil, err
	}
	return &room, nil
}
