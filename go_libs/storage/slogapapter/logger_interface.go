package slogapapter

import (
	"context"
	"fmt"
	"gorm.io/gorm/logger"
	"log/slog"
	"time"
)

type SlogAdapter struct {
	Log *slog.Logger
}

func (l *SlogAdapter) LogMode(logger.LogLevel) logger.Interface {
	return l
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
	sql, rows := fc()
	l.Log.InfoContext(ctx, "gorm", slog.String("sql", sql), slog.Int64("rows", rows), slog.Duration("elapsed", time.Since(begin)))
}
