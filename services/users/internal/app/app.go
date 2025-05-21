package app

import (
	grpcusers "crm/services/users/internal/app/grpc"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcusers.App
}

func New(logger *slog.Logger, grpcPort int, dbPath string) *App {

	grpcApp := grpcusers.New(logger, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
