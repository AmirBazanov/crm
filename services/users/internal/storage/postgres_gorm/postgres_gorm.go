package postgres_gorm

import (
	"context"
	"crm/go_libs/storage/constants"
	"crm/go_libs/storage/slogapapter"
	_ "crm/go_libs/storage/slogapapter"
	databaseusers "crm/services/users/database"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

type Storage struct {
	db     *gorm.DB
	logger *slog.Logger
}

func New(log *slog.Logger, dbUrl string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{Logger: &slogapapter.SlogAdapter{Log: log}})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return &Storage{db: db, logger: log}, nil
}

func (s *Storage) UserCreate(ctx context.Context, users *databaseusers.Users) (id uint32, err error) {
	const op = "storage.postgresgorm.SaveUser"
	result := s.db.Create(&users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrCheckConstraintViolated) {
			s.logger.Error(op, gorm.ErrCheckConstraintViolated)
			return -0, constants.ErrUserAlreadyExists
		}
		s.logger.Error(op, result.Error)
		return -0, result.Error
	}
	return users.ID, nil
}

func (s *Storage) UserByEmail(ctx context.Context, email string) (users *databaseusers.Users, err error) {
	const op = "storage.postgresgorm.UserByEmail"
	users = &databaseusers.Users{}
	result := s.db.Where("email = ?", email).First(users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			s.logger.Error(op, gorm.ErrRecordNotFound)
			return nil, constants.ErrUserNotFound
		}
		s.logger.Error(op, result.Error)
		return nil, result.Error
	}
	return users, nil
}

func (s *Storage) UserByID(ctx context.Context, id uint32) (users *databaseusers.Users, err error) {
	const op = "storage.postgresgorm.UserByID"
	users = &databaseusers.Users{}
	result := s.db.Where("id = ?", id).First(users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			s.logger.Error(op, gorm.ErrRecordNotFound)
			return nil, constants.ErrUserNotFound

		}
		s.logger.Error(op, result.Error)
		return nil, result.Error
	}
	return users, nil
}

func (s *Storage) UserByUsername(ctx context.Context, username string) (users *databaseusers.Users, err error) {
	const op = "storage.postgresgorm.UserByUsername"
	users = &databaseusers.Users{}
	result := s.db.Where("username = ?", username).First(users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			s.logger.Error(op, gorm.ErrRecordNotFound)
			return nil, constants.ErrUserNotFound
		}
		s.logger.Error(op, result.Error)
		return nil, result.Error
	}
	return users, nil
}

func (s *Storage) SearchUserByCredentials(ctx context.Context, usersCred *databaseusers.Users) (users []*databaseusers.Users, err error) {
	const op = "storage.postgresgorm.SearchUserByCredentials"
	result := s.db.Where(&usersCred).Find(&users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			s.logger.Error(op, gorm.ErrRecordNotFound)
			return nil, constants.ErrUserNotFound
		}
		s.logger.Error(op, result.Error)
		return nil, result.Error
	}
	return users, nil
}

func (s *Storage) UserUpdate(ctx context.Context, users *databaseusers.Users) (user *databaseusers.Users, err error) {
	const op = "storage.postgresgorm.UserUpdate"
	var tmp databaseusers.Users
	result := s.db.Where("id = ?", users.ID).First(&tmp)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			s.logger.Error(op, gorm.ErrRecordNotFound)
			return nil, constants.ErrUserNotFound
		}
		s.logger.Error(op, result.Error)
		return nil, result.Error
	}
	result = s.db.Updates(users)
	return users, nil
}

func (s *Storage) UserDelete(ctx context.Context, id uint32) (err error) {
	const op = "storage.postgresgorm.UserDelete"
	var tmp *databaseusers.Users
	result := s.db.Where("id = ?", id).First(&tmp)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			s.logger.Error(op, gorm.ErrRecordNotFound)
			return constants.ErrUserNotFound
		}
		s.logger.Error(op, result.Error)
		return result.Error
	}
	result = s.db.Delete(&tmp)
	return nil
}

func (s *Storage) UsersGet(ctx context.Context) (users []*databaseusers.Users, err error) {
	const op = "storage.postgresgorm.UsersGet"

	result := s.db.Find(&users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			s.logger.Error(op, gorm.ErrRecordNotFound)
			return nil, constants.ErrUserNotFound
		}
		s.logger.Error(op, result.Error)
		return nil, result.Error
	}
	return users, nil
}
