package migratorgorm

import "gorm.io/gorm"

func Migrate(db *gorm.DB, schemas ...interface{}) {
	err := db.AutoMigrate(schemas...)
	if err != nil {
		panic(err)
	}
}
