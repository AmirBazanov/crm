package migratorgorm

import (
	databaseusers "crm/services/users/database"
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

	logger.Info("Seeding countries")
	var count int64
	db.Model(&databaseusers.Countries{}).Count(&count)

	if count == 0 {
		countries := []databaseusers.Countries{
			{Code: "NIL", Name: "UNSPECIFIED"},
			{Code: "EN", Name: "England"},
			{Code: "IT", Name: "Italy"},
			{Code: "FR", Name: "France"},
			{Code: "DE", Name: "Germany"},
			{Code: "RU", Name: "RUSSIA"},
		}
		if err := db.Create(&countries).Error; err != nil {
			logger.Error("Failed to seed countries", slog.String("error", err.Error()))
		} else {
			logger.Info("Countries seeded")
		}
	} else {
		logger.Info("Countries already seeded, skipping")
	}
}
