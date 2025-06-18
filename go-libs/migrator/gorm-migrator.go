package migratorgorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

func Migrate(dbUrl string, logger *slog.Logger, schemas ...interface{}) {
	db, err := gorm.Open(postgres.Open(dbUrl))
	if err != nil {
		panic("failed to connect database")
	}
	logger.Info("Migrating database")
	err = db.AutoMigrate(schemas...)
	if err != nil {
		panic(err)
	}
	logger.Info("Database migrated successfully")
}
