package repository

import (
	"testing"

	"github.com/sandisuryadi36/sansan-dashboard/core"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	err = db.AutoMigrate(core.OrmList...)
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
