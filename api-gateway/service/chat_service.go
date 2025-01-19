package service

import (
	"context"
	"project/api-gateway/helper"
	"project/api-gateway/model"
	pbChat "project/chat-service/proto"

	"go.uber.org/zap"
)

type ChatService interface {
	SaveMessage(msg *model.Message) error
	GetRoomParticipants(roomId uint) (*pbChat.RoomParticipantsResponse, error)
	GetRoomMessages(roomId, page uint) (*pbChat.PaginatedMessagesResponse, error)
	CreateRoom(roomName string, emails []string) (*pbChat.CreateRoomResponse, error)
	AddRoomParticipant(roomId uint64, email string) (*pbChat.RoomParticipantsResponse, error)
}

type chatService struct {
	serviceUrl string
	log        *zap.Logger
}

func NewChatService(serviceUrl string, log *zap.Logger) ChatService {
	return &chatService{serviceUrl, log}
}

func (s *chatService) SaveMessage(msg *model.Message) error {
	chatConn := helper.MustConnect(s.serviceUrl)
	defer chatConn.Close()
	// log.Println(msg, "************")
	chatClient := pbChat.NewChatServiceClient(chatConn)

	req := &pbChat.SaveMessageRequest{
		RoomId:        uint64(msg.RoomId),
		SenderEmail:   msg.Sender,
		Content:       msg.Content,
		AttachmentUrl: msg.AttachmentUrl,
		ReplyTo:       uint64(msg.ReplyTo),
	}
	res, err := chatClient.SaveMessage(context.Background(), req)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	msg.Id = uint(res.MessageId)
	return nil
}

func (s *chatService) GetRoomParticipants(roomId uint) (*pbChat.RoomParticipantsResponse, error) {
	chatConn := helper.MustConnect(s.serviceUrl)
	defer chatConn.Close()

	chatClient := pbChat.NewChatServiceClient(chatConn)

	req := &pbChat.GetRoomRequest{RoomId: uint64(roomId)}
	res, err := chatClient.GetRoomParticipants(context.Background(), req)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func (s *chatService) GetRoomMessages(roomId, page uint) (*pbChat.PaginatedMessagesResponse, error) {
	chatConn := helper.MustConnect(s.serviceUrl)
	defer chatConn.Close()

	chatClient := pbChat.NewChatServiceClient(chatConn)

	req := &pbChat.GetMessagesRequest{
		RoomId: uint64(roomId),
		Limit:  10,
		Page:   uint32(page),
	}
	res, err := chatClient.GetRoomMessages(context.Background(), req)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func (s *chatService) CreateRoom(roomName string, emails []string) (*pbChat.CreateRoomResponse, error) {
	chatConn := helper.MustConnect(s.serviceUrl)
	defer chatConn.Close()

	chatClient := pbChat.NewChatServiceClient(chatConn)

	req := &pbChat.CreateRoomRequest{
		RoomName:   roomName,
		UserEmails: emails,
	}
	res, err := chatClient.CreateRoom(context.Background(), req)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func (s *chatService) AddRoomParticipant(roomId uint64, email string) (*pbChat.RoomParticipantsResponse, error) {
	chatConn := helper.MustConnect(s.serviceUrl)
	defer chatConn.Close()

	chatClient := pbChat.NewChatServiceClient(chatConn)

	req := &pbChat.AddRoomParticipantRequest{
		RoomId:    roomId,
		UserEmail: email,
	}
	res, err := chatClient.AddRoomParticipant(context.Background(), req)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	return res, nil
}
