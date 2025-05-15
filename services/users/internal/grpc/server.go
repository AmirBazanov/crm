package users

import (
	"context"
	usersv1 "crm/proto/generated/go/usersv1"

	"google.golang.org/grpc"
)

type serverAPI struct {
	usersv1.UnimplementedUserServiceServer
}

func Register(gRPC *grpc.Server) {
	usersv1.RegisterUserServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) CreateUser(ctx context.Context, req *usersv1.CreateUserRequest) (*usersv1.CreateUserResponse, error) {
	return &usersv1.CreateUserResponse{
		User: &usersv1.User{
			Id:        "1",
			Firstname: "OOO",
			Lastname:  "",
			Nickname:  "",
			Email:     "",
			Password:  "",
			Country:   5,
			CreatedAt: "",
			UpdatedAt: "",
		},
	}, nil
}

func (s *serverAPI) GetUser(ctx context.Context, req *usersv1.GetUserRequest) (*usersv1.GetUserResponse, error) {
	panic("implement me")
}

func (s *serverAPI) UpdateUser(ctx context.Context, req *usersv1.UpdateUserRequest) (*usersv1.UpdateUserResponse, error) {
	panic("implement me")
}

func (s *serverAPI) DeleteUser(ctx context.Context, req *usersv1.DeleteUserRequest) (*usersv1.DeleteUserResponse, error) {
	panic("implement me")
}

func (s *serverAPI) GetAllUsers(ctx context.Context, req *usersv1.GetUsersRequest) (*usersv1.GetUsersResponse, error) {
	panic("implement me")
}

func (s *serverAPI) SearchUsers(ctx context.Context, req *usersv1.SearchUsersRequest) (*usersv1.SearchUsersResponse, error) {
	panic("implement me")
}
