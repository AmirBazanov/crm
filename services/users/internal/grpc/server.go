package users

import (
	"buf.build/go/protovalidate"
	"context"
	"crm/go-libs/storage/constants"
	usersv2 "crm/proto/gen/go/users/v2"
	databaseusers "crm/services/users/database"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"strconv"
)

type User interface {
	Create(ctx context.Context, users *databaseusers.Users) (id string, err error)
	GetById(ctx context.Context, id string) (users *databaseusers.Users, err error)
}

type serverAPI struct {
	usersv2.UnimplementedUserServiceServer
	logger *slog.Logger
	user   User
}

func Register(gRPC *grpc.Server, logger *slog.Logger, user User) {

	usersv2.RegisterUserServiceServer(gRPC, &serverAPI{logger: logger, user: user})
}

// TODO: PROTO VALIDATION

func (s *serverAPI) CreateUser(ctx context.Context, req *usersv2.CreateUserRequest) (*usersv2.CreateUserResponse, error) {
	// TODO: IMPLEMENT PROTO VALIDATION CORRECTLY
	v, errV := protovalidate.New()
	if errV != nil {
		panic(errV)
	}
	err := v.Validate(req)
	if err != nil {
		return nil, err
	}
	op := "server.CreateUser"
	userReq := &databaseusers.Users{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Nickname:  req.Nickname,
		Email:     req.Email,
		Password:  req.Password,
		Country:   int(req.Country),
	}
	user, err := s.user.Create(ctx, userReq)
	if errors.Is(err, constants.ErrUserAlreadyExists) {
		s.logger.Error(op, err)
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	if err != nil {
		s.logger.Error(op, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &usersv2.CreateUserResponse{Id: user}, nil
}

func (s *serverAPI) GetUser(ctx context.Context, req *usersv2.GetUserRequest) (*usersv2.GetUserResponse, error) {
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
	return &usersv2.GetUserResponse{User: &usersv2.User{
		Id:        strconv.Itoa(user.ID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Country:   usersv2.Country(user.Country),
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}}, nil
}

func (s *serverAPI) UpdateUser(ctx context.Context, req *usersv2.UpdateUserRequest) (*usersv2.UpdateUserResponse, error) {
	panic("implement me")
}

func (s *serverAPI) DeleteUser(ctx context.Context, req *usersv2.DeleteUserRequest) (*usersv2.DeleteUserResponse, error) {
	panic("implement me")
}

func (s *serverAPI) GetAllUsers(ctx context.Context, req *usersv2.GetUsersRequest) (*usersv2.GetUsersResponse, error) {
	panic("implement me")
}

func (s *serverAPI) SearchUsers(ctx context.Context, req *usersv2.SearchUsersRequest) (*usersv2.SearchUsersResponse, error) {
	panic("implement me")
}
