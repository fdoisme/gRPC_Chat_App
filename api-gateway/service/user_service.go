package service

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"project/api-gateway/helper"
	"project/api-gateway/model"
	pbUser "project/user-service/proto"
)

type UserService interface {
	CreateUser(user model.User) (*pbUser.UserResponseSuccess, error)
	GetAllUsers(filter string) (*pbUser.UsersList, error)
	UpdateUser(user model.User) (*pbUser.UserResponseSuccess, error)
}

type userService struct {
	serviceUrl string
	log        *zap.Logger
}

func NewUserService(serviceUrl string, log *zap.Logger) UserService {
	return &userService{serviceUrl, log}
}

func (s *userService) CreateUser(user model.User) (*pbUser.UserResponseSuccess, error) {
	userConn := helper.MustConnect(s.serviceUrl)
	defer userConn.Close()

	userClient := pbUser.NewUserServiceClient(userConn)

	req := &pbUser.CreateUserRequest{Email: user.Email}
	res, err := userClient.CreateUser(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) GetAllUsers(filter string) (*pbUser.UsersList, error) {
	userConn := helper.MustConnect(s.serviceUrl)
	defer userConn.Close()

	userClient := pbUser.NewUserServiceClient(userConn)
	md := metadata.Pairs(
		"filter", filter,
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	req := &pbUser.Empty{}
	res, err := userClient.GetAllUsers(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) UpdateUser(user model.User) (*pbUser.UserResponseSuccess, error) {
	userConn := helper.MustConnect(s.serviceUrl)
	defer userConn.Close()

	userClient := pbUser.NewUserServiceClient(userConn)

	req := &pbUser.UpdateUserRequest{Email: user.Email, FirstName: *user.FirstName, LastName: *user.LastName}
	res, err := userClient.UpdateUser(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
