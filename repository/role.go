package repository

import (
	"context"
	"time"

	rolev1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	"github.com/sandisuryadi36/sansan-dashboard/libs/logger"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRoleList(context.Context, *rolev1.Role) ([]*rolev1.Role, error)
	GetRole(context.Context, *rolev1.Role) (*rolev1.Role, error)
	AddRole(context.Context, *rolev1.Role) (*rolev1.Role, error)
	EditRole(context.Context, *rolev1.Role) (*rolev1.Role, error)
	RemoveRole(context.Context, *rolev1.Role) (*rolev1.Role, error)
}

type gormRoleRepo struct {
	db *gorm.DB
}

func NewRoleRepository(dbMain *gorm.DB) *gormRoleRepo {
	return &gormRoleRepo{
		db: dbMain,
	}
}

func (repo *gormRoleRepo) GetRoleList(ctx context.Context, req *rolev1.Role) (res []*rolev1.Role, err error) {
	query := repo.db.Model(&rolev1.RoleORM{}).
		Where("deleted_at IS NULL").
		Order("created_at DESC")

	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		query = query.Where(&reqORM)
	}

	resORM := []*rolev1.RoleORM{}
	err = query.Find(&resORM).Error
	if err != nil {
		logger.Errorln("Fail to get role list from DB")
		return
	}

	for _, orm := range resORM {
		tmp, err := orm.ToPB(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert response to PB")
			return res, err
		}
		res = append(res, &tmp)
	}
	return
}

func (repo *gormRoleRepo) GetRole(ctx context.Context, req *rolev1.Role) (res *rolev1.Role, err error) {
	roleORM := &rolev1.RoleORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		roleORM = &reqORM
	}

	err = repo.db.Where("deleted_at IS NULL").First(roleORM, roleORM).Error
	if err != nil {
		logger.Errorln("Fail to get role list from DB")
		return
	}

	rolePB, err := roleORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return res, err
	}

	return &rolePB, err
}

func (repo *gormRoleRepo) AddRole(ctx context.Context, req *rolev1.Role) (*rolev1.Role, error) {
	roleORM := &rolev1.RoleORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		roleORM = &reqORM
	}

	err := repo.db.Create(roleORM).Error
	if err != nil {
		logger.Errorln("Fail to add role to DB")
		return nil, err
	}

	resPB, err := roleORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}
	
	return &resPB, nil
}

func (repo *gormRoleRepo) EditRole(ctx context.Context, req *rolev1.Role) (*rolev1.Role, error) {
	roleORM := &rolev1.RoleORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		roleORM = &reqORM
	}

	err := repo.db.Save(roleORM).Error
	if err != nil {
		logger.Errorln("Fail to edit role in DB")
		return nil, err
	}

	resPB, err := roleORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}
	
	return &resPB, nil
}

func (repo *gormRoleRepo) RemoveRole(ctx context.Context, req *rolev1.Role) (*rolev1.Role, error) {
	roleORM := &rolev1.RoleORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		roleORM = &reqORM
	}

	err := repo.db.Model(roleORM).Update("deleted_at", time.Now()).Error
	if err != nil {
		logger.Errorln("Fail to soft delete role from DB")
		return nil, err
	}

	resPB, err := roleORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}
	
	return &resPB, nil
}
