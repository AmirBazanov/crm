package users

import (
	"context"
	"crm/go-libs/storage/constants"
	databaseusers "crm/services/users/database"
	postgres_gorm "crm/services/users/internal/storage/postgres-gorm"
	"errors"
	"log/slog"
)

type User struct {
	logger *slog.Logger
	db     *postgres_gorm.Storage
}

type UserCreate interface {
	UserCreate(ctx context.Context, users *databaseusers.Users) (id string, err error)
}

type UserGetById interface {
	UserByID(ctx context.Context, id string) (users *databaseusers.Users, err error)
}

type UserGetByName interface {
	UserByUsername(ctx context.Context, username string) (users *databaseusers.Users, err error)
}

type UserGetByEmail interface {
	UserByEmail(ctx context.Context, email string) (users *databaseusers.Users, err error)
}

type SearchByCredentials interface {
	SearchUserByCredentials(ctx context.Context, usersCred *databaseusers.Users) (users *databaseusers.Users, err error)
}

func New(logger *slog.Logger, db *postgres_gorm.Storage) *User {
	return &User{
		logger: logger,
		db:     db,
	}
}

func (u *User) Create(ctx context.Context, users *databaseusers.Users) (id string, err error) {
	const op = "User.Create"
	u.logger.Info("creating user", op)
	resId, resErr := u.db.UserCreate(ctx, users)
	if errors.Is(resErr, constants.ErrUserAlreadyExists) {
		u.logger.Warn(op, constants.ErrUserAlreadyExists)
		return "", constants.ErrUserAlreadyExists
	}
	if resErr != nil {
		u.logger.Error(op, resErr)
		return "", resErr
	}
	return resId, nil
}

func (u *User) GetById(ctx context.Context, id string) (users *databaseusers.Users, err error) {
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
