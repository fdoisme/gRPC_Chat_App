package service

import (
	"context"
	"go.uber.org/zap"
	"project/api-gateway/helper"
	"project/api-gateway/model"
	pbAuth "project/auth-service/proto"
)

type AuthService interface {
	Register(user model.User) (*pbAuth.RegisterResponse, error)
	Login(user model.User) (*pbAuth.LoginResponse, error)
	ValidateOtp(otp model.Otp) (*string, error)
}
type authService struct {
	serviceUrl string
	log        *zap.Logger
}

func NewAuthService(serviceUrl string, log *zap.Logger) AuthService {
	return &authService{serviceUrl, log}
}

func (s *authService) Register(user model.User) (*pbAuth.RegisterResponse, error) {
	authConn := helper.MustConnect(s.serviceUrl)
	defer authConn.Close()

	authClient := pbAuth.NewAuthServiceClient(authConn)

	req := &pbAuth.RegisterRequest{Email: user.Email}
	res, err := authClient.Register(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authService) Login(user model.User) (*pbAuth.LoginResponse, error) {
	authConn := helper.MustConnect(s.serviceUrl)
	defer authConn.Close()

	authClient := pbAuth.NewAuthServiceClient(authConn)

	req := &pbAuth.LoginRequest{Email: user.Email}
	res, err := authClient.Login(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authService) ValidateOtp(otp model.Otp) (*string, error) {
	authConn := helper.MustConnect(s.serviceUrl)
	defer authConn.Close()

	authClient := pbAuth.NewAuthServiceClient(authConn)
	req := &pbAuth.ValidateOtpRequest{Id: otp.ID.String(), Otp: otp.Otp}
	res, err := authClient.ValidateOtp(context.Background(), req)
	if err != nil {
		return nil, nil
	}
	return &res.Token, nil
}
