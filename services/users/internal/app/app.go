package app

import (
	"context"
	grpcusers "crm/services/users/internal/app/grpc"
	"crm/services/users/internal/config"
	"crm/services/users/internal/services/users"
	postgresgorm "crm/services/users/internal/storage/postgres_gorm"
	"crm/services/users/pkg/redis"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcusers.App
}

func New(logger *slog.Logger, grpcPort int, dbUrl string, redisCfg config.RedisConfig, ctx context.Context) *App {
	storage, err := postgresgorm.New(logger, dbUrl)
	cache, errRed := redis.NewClient(logger, redisCfg.Addr, redisCfg.Password, redisCfg.DB, ctx)
	if errRed != nil {
		logger.Error("Cannot connect to redis", errRed)
	}
	if err != nil {
		panic(err)
	}
	userService := users.New(logger, storage, cache)
	grpcApp := grpcusers.New(logger, grpcPort, userService)

	return &App{
		GRPCSrv: grpcApp,
	}
}
