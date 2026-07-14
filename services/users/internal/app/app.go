package app

import (
	"context"
	grpcusers "crm/services/users/internal/app/grpc"
	"crm/services/users/internal/config"
	"crm/services/users/internal/events/kafka/handlers"
	"crm/services/users/internal/kafka"
	"crm/services/users/internal/services/users"
	postgresgorm "crm/services/users/internal/storage/postgres_gorm"
	"crm/services/users/pkg/redis"
	"log/slog"
)

type App struct {
	GRPCSrv  *grpcusers.App
	Consumer *kafka.Consumer
}

func New(logger *slog.Logger, grpcPort int, dbUrl string, redisCfg config.RedisConfig, kafkaCfg config.KafkaConfig) *App {
	storage, err := postgresgorm.New(logger, dbUrl)
	cache, errRed := redis.NewClient(logger, redisCfg.Addr, redisCfg.Password, redisCfg.DB)
	if errRed != nil {
		logger.Error("Cannot create redis client " + errRed.Error())
	} else {
		errRed = cache.Ping(context.Background())
		if errRed != nil {
			logger.Error("Cannot connect to redis " + errRed.Error())
		}
	}
	if err != nil {
		panic(err)
	}
	userService := users.New(logger, storage)
	grpcApp := grpcusers.New(logger, grpcPort, userService, cache)
	router := kafka.NewRouter()
	router.Register("create-user", handlers.NewUserCreatedHandler(userService))
	consumer := kafka.New(kafkaCfg.Brokers, kafkaCfg.Topic, kafkaCfg.GroupID, logger, router)

	return &App{
		GRPCSrv:  grpcApp,
		Consumer: consumer,
	}
}
