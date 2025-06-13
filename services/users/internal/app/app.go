package app

import (
	grpcusers "crm/services/users/internal/app/grpc"
	"crm/services/users/internal/services/users"
	postgresgorm "crm/services/users/internal/storage/postgres-gorm"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcusers.App
}

func New(logger *slog.Logger, grpcPort int, dbUrl string) *App {
	storage, err := postgresgorm.New(logger, dbUrl)
	if err != nil {
		panic(err)
	}
	userService := users.New(logger, storage)
	grpcApp := grpcusers.New(logger, grpcPort, userService)

	return &App{
		GRPCSrv: grpcApp,
	}
}
