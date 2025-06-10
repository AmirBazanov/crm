package users

import (
	"context"
	usersv1 "crm/proto/gen/go/users/v1"
	"google.golang.org/grpc"
	"log/slog"
)

type serverAPI struct {
	usersv1.UnimplementedUserServiceServer
	logger *slog.Logger
}

func Register(gRPC *grpc.Server, logger *slog.Logger) {

	usersv1.RegisterUserServiceServer(gRPC, &serverAPI{logger: logger})
}

func (s *serverAPI) CreateUser(ctx context.Context, req *usersv1.CreateUserRequest) (*usersv1.CreateUserResponse, error) {
	panic("implement me")
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
