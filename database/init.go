package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	dbFile := "C:/Users/Administrator/GolandProjects/come2/data/come2.db"

	gdb, err := gorm.Open(sqlite.Open(dbFile),
		&gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Info),
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		return err
	}

	DB = gdb
	return nil
}
