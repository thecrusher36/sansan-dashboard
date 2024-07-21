package repository

import (
	"testing"

	"github.com/sandisuryadi36/sansan-dashboard/core"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	err = db.AutoMigrate(core.OrmList...)
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
