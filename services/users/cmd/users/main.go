package main

import (
	gologger "crm/libs/go-logger"
	"crm/services/users/internal/config"
	"fmt"
	"log/slog"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	logger := setupLogger(cfg)
	logger.Error("Logger initialized")
}

func setupLogger(config *config.Config) *slog.Logger {
	logCfg := config.Logger
	gologger.InitLogger(logCfg.Service, logCfg.LogLevel, logCfg.LogFile)
	return gologger.GetLogger()
}
