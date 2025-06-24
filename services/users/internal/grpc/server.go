package users

import (
	"context"
	_map "crm/go_libs/helpers/map"
	"crm/go_libs/helpers/pointercheck"

	"crm/go_libs/storage/constants"
	usersv3 "crm/proto/gen/go/users/v3"
	databaseusers "crm/services/users/database"
	"crm/services/users/internal/tools/convert"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"log/slog"
)

type User interface {
	Create(ctx context.Context, users *databaseusers.Users) (id uint32, err error)
	GetById(ctx context.Context, id uint32) (users *databaseusers.Users, err error)
	GetByUsername(ctx context.Context, username string) (users *databaseusers.Users, err error)
	GetByEmail(ctx context.Context, email string) (users *databaseusers.Users, err error)
	Update(ctx context.Context, users *databaseusers.Users) (user *databaseusers.Users, err error)
	Delete(ctx context.Context, id uint32) (err error)
	Search(ctx context.Context, usersCred *databaseusers.Users) (users []*databaseusers.Users, err error)
	Users(ctx context.Context) (users []*databaseusers.Users, err error)
}

type serverAPI struct {
	usersv3.UnimplementedUserServiceServer
	logger *slog.Logger
	user   User
}

func Register(gRPC *grpc.Server, logger *slog.Logger, user User) {

	usersv3.RegisterUserServiceServer(gRPC, &serverAPI{logger: logger, user: user})
}

func (s *serverAPI) CreateUser(ctx context.Context, req *usersv3.CreateUserRequest) (*usersv3.CreateUserResponse, error) {
	op := "server.CreateUser"
	userReq := &databaseusers.Users{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Nickname:  req.Nickname,
		Email:     req.Email,
		Password:  req.Password,
		Country:   int(req.Country) + 1,
	}
	//TODO: FIX Unspecified country
	user, err := s.user.Create(ctx, userReq)
	if errors.Is(err, constants.ErrUserAlreadyExists) {
		s.logger.Error(op, err)
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	if err != nil {
		s.logger.Error(op, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &usersv3.CreateUserResponse{Id: user}, nil
}

func (s *serverAPI) GetUser(ctx context.Context, req *usersv3.GetUserRequest) (*usersv3.GetUserResponse, error) {
	op := "server.GetUser"
	user, err := s.user.GetById(ctx, req.Id)
	if errors.Is(err, constants.ErrUserNotFound) {
		s.logger.Error(op, err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		s.logger.Error(op, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &usersv3.GetUserResponse{User: convert.UserDbtoGrpcUser(user)}, nil
}

func (s *serverAPI) UpdateUser(ctx context.Context, req *usersv3.UpdateUserRequest) (*usersv3.UpdateUserResponse, error) {
	op := "server.UpdateUser"
	userReq := &databaseusers.Users{
		ID:        req.Id,
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Nickname:  req.Nickname,
		Email:     req.Email,
		Country:   int(req.Country),
	}
	user, err := s.user.Update(ctx, userReq)
	if errors.Is(err, constants.ErrUserNotFound) {
		s.logger.Error(op, err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		s.logger.Error(op, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &usersv3.UpdateUserResponse{User: convert.UserDbtoGrpcUser(user)}, nil
}

func (s *serverAPI) GetUsers(ctx context.Context, req *usersv3.GetUsersRequest) (*usersv3.GetUsersResponse, error) {
	op := "server.GetAllUsers"
	users, err := s.user.Users(ctx)
	if errors.Is(err, constants.ErrUserNotFound) {
		s.logger.Error(op, err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		s.logger.Error(op, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &usersv3.GetUsersResponse{Users: _map.Map(users, convert.UserDbtoGrpcUser)}, nil
}

func (s *serverAPI) Search(ctx context.Context, req *usersv3.SearchUsersRequest) (*usersv3.SearchUsersResponse, error) {
	userSear := databaseusers.Users{
		FirstName: pointercheck.DerefOrDefault(req.Firstname, ""),
		LastName:  pointercheck.DerefOrDefault(req.Lastname, ""),
		Nickname:  pointercheck.DerefOrDefault(req.Nickname, ""),
		Email:     pointercheck.DerefOrDefault(req.Email, ""),
		Country:   int(pointercheck.DerefOrDefault(req.Country, 0)),
	}
	op := "server.SearchUsers"
	user, err := s.user.Search(ctx, &userSear)
	if errors.Is(err, constants.ErrUserNotFound) {
		s.logger.Error(op, err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		s.logger.Error(op, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &usersv3.SearchUsersResponse{Users: _map.Map(user, convert.UserDbtoGrpcUser)}, nil
}

func (s *serverAPI) DeleteUser(ctx context.Context, req *usersv3.DeleteUserRequest) (*usersv3.DeleteUserResponse, error) {
	op := "server.DeleteUser"
	err := s.user.Delete(ctx, req.Id)
	if errors.Is(err, constants.ErrUserNotFound) {
		s.logger.Error(op, err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		s.logger.Error(op, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &usersv3.DeleteUserResponse{}, nil
}
