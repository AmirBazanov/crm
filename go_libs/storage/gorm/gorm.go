package gormstorage_slogadapter

import (
	"crm/go_libs/storage/slogapapter"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

type GORM struct {
	DB *gorm.DB
}

func New(log *slog.Logger, dbUrl string) *GORM {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{Logger: &slogapapter.SlogAdapter{Log: log}})
	if err != nil {
		panic(err)
	}
	return &GORM{DB: db}
}
