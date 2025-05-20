package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	once   sync.Once
	logger *slog.Logger
)

// InitLogger инициализирует глобальный логгер.
func InitLogger(service string, logLevel string, logFile string) *slog.Logger {
	once.Do(func() {
		level := getLogLevel(logLevel)
		handler := createHandler(level, service, logFile)

		logger = slog.New(handler)
	})
	return logger
}

func getLogLevel(logLevel string) slog.Level {
	envLogLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if envLogLevel != "" {
		logLevel = envLogLevel
	}
	switch logLevel {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

type CustomHandler struct {
	Level   slog.Level
	Service string
	Writer  io.Writer
}

func (c *CustomHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= c.Level
}

func (c *CustomHandler) Handle(_ context.Context, record slog.Record) error {
	timestamp := record.Time.UTC().Format(time.RFC3339Nano)
	level := strings.ToUpper(record.Level.String())

	// Формируем основное сообщение
	message := record.Message

	// Считываем ключи и значения
	var attrs []string
	record.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, fmt.Sprintf("%s=%v", a.Key, a.Value.Any()))
		return true
	})

	var extra string
	if len(attrs) > 0 {
		extra = " " + strings.Join(attrs, " ")
	}

	// Печатаем в нужном формате
	_, err := fmt.Fprintf(c.Writer, "%s [%s] [%s] %s%s\n", timestamp, level, c.Service, message, extra)
	return err
}

func (c *CustomHandler) WithAttrs([]slog.Attr) slog.Handler {
	return c
}

func (c *CustomHandler) WithGroup(string) slog.Handler {
	return c
}

func createHandler(level slog.Level, service string, logFile string) slog.Handler {
	if logFile == "" {
		logFile = os.Getenv("LOG_FILE")
	}
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("could not open log file: " + err.Error())
	}

	writer := io.MultiWriter(os.Stdout, file)

	return &CustomHandler{
		Level:   level,
		Service: service,
		Writer:  writer,
	}
}

// GetLogger возвращает инициализированный логгер (если InitLogger уже был вызван).
func GetLogger() *slog.Logger {
	if logger == nil {
		panic("Logger is not initialized. Call InitLogger first.")
	}
	return logger
}
