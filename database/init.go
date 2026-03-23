package database

import (
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	rootDir := filepath.FromSlash("C:/Users/Administrator/GolandProjects/come2")
	dbDir := filepath.Join(rootDir, "data")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		return err
	}

	dbFile := filepath.Join(dbDir, "come2.db")

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

// GetDB returns the initialized *gorm.DB (may be nil if InitDB not called or failed).
func GetDB() *gorm.DB {
	return DB
}
