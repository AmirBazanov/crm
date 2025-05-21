package gormstorage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

type GORM struct {
	DB *gorm.DB
}

func New(log *slog.Logger, dbUrl string) *GORM {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{Logger: &SlogAdapter{log: log}})
	if err != nil {
		panic(err)
	}
	return &GORM{DB: db}
}
