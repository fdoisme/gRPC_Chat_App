package handler

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project/chat-service/helper"
	"project/chat-service/model"
	pb "project/chat-service/proto"
	"project/chat-service/service"
)

type ChatHandler struct {
	Service service.Service
	Logger  *zap.Logger
	pb.UnimplementedChatServiceServer
}

func NewChatHandler(service service.Service, logger *zap.Logger) *ChatHandler {
	return &ChatHandler{
		Service: service,
		Logger:  logger,
	}
}

func (h *ChatHandler) AddRoomParticipant(ctx context.Context, req *pb.AddRoomParticipantRequest) (*pb.RoomParticipantsResponse, error) {
	h.Logger.Info("Received AddRoomParticipant request",
		zap.Uint64("roomId", req.GetRoomId()),
		zap.String("userEmail", req.UserEmail),
	)

	// Fetch room details to ensure it exists
	room, err := h.Service.ChatService.GetRoomDetails(uint(req.GetRoomId()))
	if err != nil {
		h.Logger.Error("Error fetching room details", zap.Error(err))
		return nil, status.Errorf(codes.NotFound, "room not found: %v", err)
	}

	// Check if the user is already a participant in the room
	existingParticipants, err := h.Service.ChatService.GetRoomParticipants(uint(req.GetRoomId()))
	if err != nil {
		h.Logger.Error("Error fetching existing participants", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to fetch participants: %v", err)
	}

	for _, participant := range existingParticipants {
		if participant.UserEmail == req.UserEmail {
			h.Logger.Warn("User already a participant", zap.String("email", req.UserEmail))
			return nil, status.Errorf(codes.AlreadyExists, "user already a participant in the room")
		}
	}

	// Add the new participant to the room
	newParticipant := &model.RoomParticipant{
		RoomID:    uint(req.GetRoomId()),
		UserEmail: req.UserEmail,
	}
	if err := h.Service.ChatService.CreateRoomParticipant(newParticipant); err != nil {
		h.Logger.Error("Error adding participant to room", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to add participant: %v", err)
	}

	h.Logger.Info("Participant added successfully",
		zap.Uint64("roomId", req.GetRoomId()),
		zap.String("userEmail", req.UserEmail),
	)

	// Fetch updated participants list
	updatedParticipants, err := h.Service.ChatService.GetRoomParticipants(uint(req.GetRoomId()))
	if err != nil {
		h.Logger.Error("Error fetching updated participants", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to fetch updated participants: %v", err)
	}

	// Prepare response with updated participants
	users := make([]*pb.User, len(updatedParticipants))
	for i, p := range updatedParticipants {
		users[i] = &pb.User{
			UserId:   uint64(p.UserID),
			Username: p.UserEmail,
		}
	}

	return &pb.RoomParticipantsResponse{
		RoomId:   uint64(room.ID),
		RoomName: room.Name,
		Users:    users,
	}, nil
}

func (h *ChatHandler) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	h.Logger.Info("CreateRoom request received", zap.String("roomName", req.RoomName))

	room := &model.Room{
		Name: req.RoomName,
	}

	if err := h.Service.ChatService.CreateRoom(room); err != nil {
		h.Logger.Error("Failed to create room", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to create room")
	}

	for _, email := range req.UserEmails {
		participant := &model.RoomParticipant{
			RoomID:    room.ID,
			UserEmail: email,
		}
		if err := h.Service.ChatService.CreateRoomParticipant(participant); err != nil {
			h.Logger.Error("Failed to add user", zap.String("email", email), zap.Error(err))
		}
	}

	return &pb.CreateRoomResponse{
		RoomId:   uint64(room.ID),
		RoomName: room.Name,
	}, nil
}

func (h *ChatHandler) SaveMessage(ctx context.Context, req *pb.SaveMessageRequest) (*pb.SaveMessageResponse, error) {
	h.Logger.Info("SaveMessage request", zap.Uint64("roomId", req.RoomId))

	message := &model.Message{
		RoomID:        uint(req.RoomId),
		SenderEmail:   req.SenderEmail,
		Content:       req.Content,
		AttachmentURL: helper.Ptr(req.AttachmentUrl),
		ReplyTo:       helper.Ptr(uint(req.ReplyTo)),
	}

	if err := h.Service.ChatService.SaveMessage(message); err != nil {
		h.Logger.Error("Failed to save message", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to save message")
	}

	return &pb.SaveMessageResponse{
		MessageId: uint64(message.ID),
		CreatedAt: message.CreatedAt.UTC().String(),
	}, nil
}

func (h *ChatHandler) GetRoomParticipants(ctx context.Context, req *pb.GetRoomRequest) (*pb.RoomParticipantsResponse, error) {
	h.Logger.Info("Received GetRoomParticipants request", zap.Uint64("roomId", req.RoomId))

	participants, err := h.Service.ChatService.GetRoomParticipants(uint(req.RoomId))
	if err != nil {
		h.Logger.Error("Error fetching room participants", zap.Uint64("roomId", req.RoomId), zap.Error(err))
		return nil, err
	}

	h.Logger.Info("Successfully fetched participants", zap.Int("numParticipants", len(participants)))

	users := make([]*pb.User, len(participants))
	for i, p := range participants {
		users[i] = &pb.User{
			UserId:   uint64(p.UserID),
			Username: p.UserEmail,
		}
	}

	room, err := h.Service.ChatService.GetRoomDetails(uint(req.RoomId))
	if err != nil {
		h.Logger.Error("Error fetching room details", zap.Uint64("roomId", req.RoomId), zap.Error(err))
		return nil, err
	}

	h.Logger.Info("Successfully fetched room details", zap.Int64("roomId", int64(room.ID)))

	return &pb.RoomParticipantsResponse{
		RoomId:   uint64(room.ID),
		RoomName: room.Name,
		Users:    users,
	}, nil
}

func (h *ChatHandler) GetRoomMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.PaginatedMessagesResponse, error) {
	h.Logger.Info("Received GetRoomMessages request", zap.Uint64("roomId", req.RoomId), zap.Int("limit", int(req.Limit)), zap.Int("page", int(req.Page)))

	pagination, err := h.Service.ChatService.GetRoomMessages(uint(req.RoomId), int(req.Limit), int(req.Page))
	if err != nil {
		h.Logger.Error("Error fetching room messages", zap.Uint64("roomId", req.RoomId), zap.Int("limit", int(req.Limit)), zap.Int("page", int(req.Page)), zap.Error(err))
		return nil, err
	}

	room, err := h.Service.ChatService.GetRoomDetails(uint(req.RoomId))
	if err != nil {
		h.Logger.Error("Error fetching room details", zap.Uint64("roomId", req.RoomId), zap.Error(err))
		return nil, err
	}

	h.Logger.Info("Successfully fetched room details", zap.Int64("roomId", int64(room.ID)))

	var msgs []*pb.Message
	for _, m := range pagination.Messages {
		var attachmentURL string
		if m.AttachmentURL != nil {
			attachmentURL = *m.AttachmentURL
		} else {
			attachmentURL = ""
		}

		var replyTo int64
		if m.ReplyTo != nil {
			replyTo = int64(*m.ReplyTo)
		} else {
			replyTo = 0
		}

		var readAt string
		if m.ReadAt != nil {
			readAt = m.ReadAt.String()
		} else {
			readAt = ""
		}

		data := &pb.Message{
			MessageId:     uint64(m.ID),
			SenderEmail:   m.SenderEmail,
			Content:       m.Content,
			AttachmentUrl: attachmentURL,
			ReplyTo:       uint64(replyTo),
			SentAt:        m.CreatedAt.String(),
			ReadAt:        readAt,
		}

		msgs = append(msgs, data)
	}

	return &pb.PaginatedMessagesResponse{
		RoomId:   uint64(room.ID),
		RoomName: room.Name,
		Messages: msgs,
		Pagination: &pb.Pagination{
			Page:       uint32(pagination.Page),
			Limit:      uint32(pagination.Limit),
			TotalPages: uint32(pagination.TotalPages),
			TotalItems: uint32(pagination.TotalItems),
		},
	}, nil
}
