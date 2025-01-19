package service

import (
	"context"
	"project/user-service/model"
	pb "project/user-service/proto"
	"project/user-service/repository"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type UserService struct {
	repo repository.Repository
	log  *zap.Logger
	pb.UnimplementedUserServiceServer
}

func NewUserService(repo repository.Repository, log *zap.Logger) *UserService {
	return &UserService{repo: repo, log: log}
}

func (s *UserService) GetAllUsers(ctx context.Context, req *pb.Empty) (*pb.UsersList, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	var filter bool
	isOnline, ok := md["filter"]
	if ok && isOnline[0] != "" {
		// filter = bool(isOnline[0])
		// fmt.Printf("%+v\n", md)
		// fmt.Println(filter, len(isOnline))
		filter = true
	}
	users, err := s.repo.User.GetAllUsers(filter)
	if err != nil {
		return nil, err
	}
	var usersPb []*pb.User
	for _, user := range users {
		usersPb = append(usersPb, &pb.User{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			IsOnline:  user.IsOnline,
		})

	}
	return &pb.UsersList{Users: usersPb}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponseSuccess, error) {
	var user model.User
	user.Email = req.Email
	err := s.repo.User.Insert(&user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponseSuccess{Message: "Create User Succes"}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponseSuccess, error) {
	user := model.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	err := s.repo.User.UpdateProfile(&user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponseSuccess{Message: "Update Profile Success"}, nil
}
