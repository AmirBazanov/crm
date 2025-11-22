package slogapapter

import (
	"context"
	"fmt"
	"gorm.io/gorm/logger"
	"log/slog"
	"strings"
	"time"
)

type SlogAdapter struct {
	Log   *slog.Logger
	Level logger.LogLevel
}

func (l *SlogAdapter) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (l *SlogAdapter) shouldLog(level logger.LogLevel) bool {
	return level <= l.Level
}

func (l *SlogAdapter) Info(ctx context.Context, s string, args ...interface{}) {
	l.Log.InfoContext(ctx, fmt.Sprintf(s, args...))
}

func (l *SlogAdapter) Warn(ctx context.Context, s string, args ...interface{}) {
	l.Log.WarnContext(ctx, fmt.Sprintf(s, args...))
}

func (l *SlogAdapter) Error(ctx context.Context, s string, args ...interface{}) {
	l.Log.ErrorContext(ctx, fmt.Sprintf(s, args...))
}

func (l *SlogAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), _ error) {
	sql, row := fc()
	query := ""
	if strings.Contains(sql, "INSERT INTO") {
		query = "INSERT"
	} else if strings.Contains(sql, "SELECT") {
		query = "SELECT"
	} else if strings.Contains(sql, "UPDATE") {
		query = "UPDATE"
	} else if strings.Contains(sql, "DELETE FROM") {
		query = "DELETE"
	}

	l.Log.InfoContext(ctx, "gorm", slog.String("sql", query), slog.Int64("rows", row), slog.Duration("elapsed", time.Since(begin)))

}
