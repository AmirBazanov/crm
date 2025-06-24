package migratorgorm

import (
	databaseusers "crm/services/users/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

func Migrate(dbUrl string, logger *slog.Logger, schemas ...interface{}) {
	countries := []*databaseusers.Countries{
		{Code: "NIL", Name: "UNSPECIFIED"},
		{Code: "EN", Name: "England"},
		{Code: "IT", Name: "Italy"},
		{Code: "FR", Name: "France"},
		{Code: "DE", Name: "Germany"},
		{Code: "RU", Name: "RUSSIA"},
	}
	db, err := gorm.Open(postgres.Open(dbUrl))
	if err != nil {
		panic("failed to connect database")
	}
	logger.Info("Migrating database")
	err = db.AutoMigrate(schemas...)
	err = db.AutoMigrate(countries)
	if err != nil {
		panic(err)
	}
	logger.Info("Database migrated successfully")
}
