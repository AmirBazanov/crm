package gormhealth

import (
	"database/sql"
	"gorm.io/gorm"
	"log/slog"
)

func GormHealthCheck(db *gorm.DB, logger *slog.Logger) (ok string, err error) {
	con, err := db.DB()
	defer func(con *sql.DB) {
		err := con.Close()
		if err != nil {
			logger.Error("Failed to close DB connection", err)
			return
		}
	}(con)
	if err != nil {
		logger.Error("Error connecting to database", err.Error())
		return "Error connecting to database", err
	}
	err = con.Ping()
	if err != nil {
		logger.Error("Error pinging database", err.Error())
		return "Error pinging database", err
	}
	return "Successfully connected to database", err

}
