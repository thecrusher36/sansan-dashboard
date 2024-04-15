package repository

import (
	"context"
	"time"

	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	"github.com/sandisuryadi36/sansan-dashboard/libs/logger"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserList(context.Context, *userv1.User) ([]*userv1.User, error)
	GetUser(context.Context, *userv1.User) (*userv1.User, error)
	AddUser(context.Context, *userv1.User) (*userv1.User, error)
	EditUser(context.Context, *userv1.User) (*userv1.User, error)
	RemoveUser(context.Context, *userv1.User) (*userv1.User, error)
}

type GormUserRepo struct {
	db *gorm.DB
}

func NewUserRepository(dbMain *gorm.DB) *GormUserRepo {
	return &GormUserRepo{
		db: dbMain,
	}
}

func (repo *GormUserRepo) GetUserList(ctx context.Context, req *userv1.User) ([]*userv1.User, error) {
	usersORM := []*userv1.UserORM{}
	query := repo.db.Model(&userv1.UserORM{}).
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
	err := query.Find(&usersORM).Error
	if err != nil {
		logger.Errorln("Fail to get user list from DB")
		return nil, err
	}

	users := make([]*userv1.User, len(usersORM))
	for i, userORM := range usersORM {
		userPB, err := userORM.ToPB(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert response to PB")
			return nil, err
		}
		users[i] = &userPB
	}

	return users, nil
}

func (repo *GormUserRepo) GetUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
	userORM := &userv1.UserORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		userORM = &reqORM
	}

	err := repo.db.Where("deleted_at IS NULL").First(userORM, userORM).Error
	if err != nil {
		logger.Errorln("Fail to get user from DB")
		return nil, err
	}

	userPB, err := userORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}
	
	return &userPB, nil
}

func (repo *GormUserRepo) AddUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
	userORM := &userv1.UserORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		userORM = &reqORM
	}

	err := repo.db.Create(userORM).Error
	if err != nil {
		logger.Errorln("Fail to add user to DB")
		return nil, err
	}

	resPB, err := userORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}
	
	return &resPB, nil
}

func (repo *GormUserRepo) EditUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
	userORM := &userv1.UserORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		userORM = &reqORM
	}

	err := repo.db.Save(userORM).Error
	if err != nil {
		logger.Errorln("Fail to edit user in DB")
		return nil, err
	}

	resPB, err := userORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}
	
	return &resPB, nil
}

func (repo *GormUserRepo) RemoveUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
	userORM := &userv1.UserORM{}
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, err
		}
		userORM = &reqORM
	}

	err := repo.db.Model(userORM).Update("deleted_at", time.Now()).Error
	if err != nil {
		logger.Errorln("Fail to soft delete user from DB")
		return nil, err
	}

	resPB, err := userORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}
	
	return &resPB, nil
}
