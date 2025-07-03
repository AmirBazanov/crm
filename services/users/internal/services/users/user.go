package users

import (
	"context"
	"crm/go_libs/storage/constants"
	databaseusers "crm/services/users/database"
	"errors"
	"log/slog"
)

type User struct {
	logger *slog.Logger
	db     Storage
}

type Storage interface {
	UserByID(ctx context.Context, id uint32) (users *databaseusers.Users, err error)
	UserByUsername(ctx context.Context, username string) (users *databaseusers.Users, err error)
	UserByEmail(ctx context.Context, email string) (users *databaseusers.Users, err error)
	UsersGet(ctx context.Context) (users []*databaseusers.Users, err error)
	UserCreate(ctx context.Context, users *databaseusers.Users) (id uint32, err error)
	UserUpdate(ctx context.Context, users *databaseusers.Users) (user *databaseusers.Users, err error)
	UserDelete(ctx context.Context, id uint32) (err error)
	SearchUserByCredentials(ctx context.Context, usersCred *databaseusers.Users) (users []*databaseusers.Users, err error)
}

func New(logger *slog.Logger, db Storage) *User {
	return &User{
		logger: logger,
		db:     db,
	}
}

func (u *User) Create(ctx context.Context, users *databaseusers.Users) (id uint32, err error) {
	const op = "User.Create"
	u.logger.Info("creating user", op, users.ID)
	resId, resErr := u.db.UserCreate(ctx, users)
	if errors.Is(resErr, constants.ErrUserAlreadyExists) {
		u.logger.Warn(op, constants.ErrUserAlreadyExists)
		return -0, constants.ErrUserAlreadyExists
	}
	if resErr != nil {
		u.logger.Error(op, resErr)
		return -0, resErr
	}
	return resId, nil
}

func (u *User) GetById(ctx context.Context, id uint32) (users *databaseusers.Users, err error) {
	const op = "User.GetById"
	u.logger.Info("getting user", op, id)
	resUser, resErr := u.db.UserByID(ctx, id)
	if errors.Is(resErr, constants.ErrUserNotFound) {
		u.logger.Warn(op, constants.ErrUserNotFound)
		return nil, constants.ErrUserNotFound
	}
	if resErr != nil {
		u.logger.Error(op, resErr)
		return nil, resErr
	}
	return resUser, nil
}

func (u *User) GetByUsername(ctx context.Context, username string) (users *databaseusers.Users, err error) {
	const op = "User.GetByUsername"
	u.logger.Info("getting user by username", op, username)
	resUser, resErr := u.db.UserByUsername(ctx, username)
	if errors.Is(resErr, constants.ErrUserNotFound) {
		u.logger.Warn(op, constants.ErrUserNotFound)
		return nil, constants.ErrUserNotFound
	}
	if resErr != nil {
		u.logger.Error(op, resErr)
		return nil, resErr
	}
	return resUser, nil
}

func (u *User) GetByEmail(ctx context.Context, email string) (users *databaseusers.Users, err error) {
	const op = "User.GetByEmail"
	u.logger.Info("getting user by email", op, email)
	resUser, resErr := u.db.UserByEmail(ctx, email)
	if errors.Is(resErr, constants.ErrUserNotFound) {
		u.logger.Warn(op, constants.ErrUserNotFound)
		return nil, constants.ErrUserNotFound
	}
	if resErr != nil {
		u.logger.Error(op, resErr)
		return nil, resErr
	}
	return resUser, nil
}

func (u *User) Update(ctx context.Context, users *databaseusers.Users) (user *databaseusers.Users, err error) {
	const op = "User.Update"
	u.logger.Info("updating user", op, users)
	resUser, resErr := u.db.UserUpdate(ctx, users)
	if errors.Is(resErr, constants.ErrUserNotFound) {
		u.logger.Warn(op, constants.ErrUserNotFound)
		return nil, constants.ErrUserNotFound
	}
	if resErr != nil {
		u.logger.Error(op, resErr)
		return nil, resErr
	}
	return resUser, nil
}

func (u *User) Delete(ctx context.Context, id uint32) (err error) {
	const op = "User.Delete"
	u.logger.Info("deleting user", op, id)
	resErr := u.db.UserDelete(ctx, id)
	if errors.Is(resErr, constants.ErrUserNotFound) {
		u.logger.Warn(op, constants.ErrUserNotFound)
		return constants.ErrUserNotFound

	}
	if resErr != nil {
		u.logger.Error(op, resErr)
		return resErr
	}
	return nil
}

func (u *User) Search(ctx context.Context, usersCred *databaseusers.Users) (users []*databaseusers.Users, err error) {
	const op = "User.Search"
	u.logger.Info("searching users", op, usersCred)
	resU, resErr := u.db.SearchUserByCredentials(ctx, usersCred)
	if errors.Is(resErr, constants.ErrUserNotFound) {
		u.logger.Warn(op, constants.ErrUserNotFound)
		return nil, constants.ErrUserNotFound
	}
	if resErr != nil {
		u.logger.Error(op, resErr)
		return nil, resErr
	}
	return resU, nil
}

func (u *User) Users(ctx context.Context) (users []*databaseusers.Users, err error) {
	const op = "User.Users"
	u.logger.Info("getting users", op)
	resU, resErr := u.db.UsersGet(ctx)
	if errors.Is(resErr, constants.ErrUserNotFound) {
		u.logger.Warn(op, constants.ErrUserNotFound)
		return nil, constants.ErrUserNotFound
	}
	if resErr != nil {
		u.logger.Error(op, resErr)
	}
	return resU, nil
}
