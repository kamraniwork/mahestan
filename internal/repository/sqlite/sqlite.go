package sqlite

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mahestan/conf"
)

func ConnectToSqlite(cfg *conf.AppConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s", cfg.Database.Sqlite.Name)

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, err
}
