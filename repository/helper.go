package repository

import (
	"testing"

	rolev1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	err = db.AutoMigrate(
		&rolev1.RoleORM{},
		&userv1.UserORM{},
	)
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
