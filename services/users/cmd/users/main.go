package main

import (
	gologger "crm/go-libs/logger"
	migratorgorm "crm/go-libs/migrator"
	gormstorage "crm/go-libs/storage/gorm"
	databaseusers "crm/services/users/database"
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
	application := app.New(logger, cfg.GRPC.Port, cfg.DbUrl)
	db := gormstorage.New(logger, cfg.DbUrl)
	migratorgorm.Migrate(db.DB, &databaseusers.Countries{}, &databaseusers.Users{})
	go application.GRPCSrv.MustRun()
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
	logger.Info("Application stopped", "signal", sign)
}
