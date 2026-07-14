package main

import (
	"context"
	gologger "crm/go_libs/logger"
	"crm/services/users/internal/app"
	"crm/services/users/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	logger := setupLogger(cfg)
	// migratorgorm.Migrate(cfg.DbUrl, logger, &databaseusers.Countries{}, &databaseusers.Users{})
	application := app.New(logger, cfg.GRPC.Port, cfg.DbUrl, cfg.Redis, cfg.Kafka)
	go application.GRPCSrv.MustRun()
	go application.Consumer.Start(context.Background())
	GrpcStop(application, logger)
}

func setupLogger(config *config.Config) *slog.Logger {
	logCfg := config.Logger
	gologger.InitLogger(logCfg.Service, logCfg.LogLevel, logCfg.LogFile)
	return gologger.GetLogger()
}

func GrpcStop(app *app.App, logger *slog.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	sign := <-stop
	app.GRPCSrv.Stop()
	if err := app.Consumer.Stop(); err != nil {
		logger.Error("failed to stop consumer", "error", err)
	}
	logger.Info("Application stopped", "signal", sign)
}
