package repository

import (
	"context"

	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserList(context.Context) ([]*userv1.User, error)
}

type GormUserRepo struct {
	db *gorm.DB
}

func NewUserRepository(dbMain *gorm.DB) *GormUserRepo {
	return &GormUserRepo{
		db: dbMain,
	}
}
