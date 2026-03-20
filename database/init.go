package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *DB

func InitDB() error {
	DB, err = gorm.Open(sqlite.Open("jdbc:sqlite:C:/Users/Administrator/GolandProjects/come2/data"),
		&gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Info),
			SkipDefaultTransaction: true,
		},
	)
	return err
}
