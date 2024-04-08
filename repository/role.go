package repository

import (
	"context"

	rolev1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRoleList(context.Context) ([]*rolev1.Role, error)
}

type GormRoleRepo struct {
	db *gorm.DB
}

func NewRoleRepository(dbMain *gorm.DB) *GormRoleRepo {
	return &GormRoleRepo{
		db: dbMain,
	}
}

func (repo * GormRoleRepo) GetRoleList(ctx context.Context) ([]*rolev1.Role, error) {
	return rolev1.DefaultListRole(ctx, repo.db)
}