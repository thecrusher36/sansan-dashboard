package repository

import (
	"context"
	"strings"
	"time"

	commonv1 "github.com/sandisuryadi36/sansan-dashboard/gen/common/v1"
	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	"github.com/sandisuryadi36/sansan-dashboard/libs/logger"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserList(context.Context, *userv1.User, *commonv1.StandardQuery) ([]*userv1.User, *commonv1.StandardPaginationResponse, error)
	GetUser(context.Context, *userv1.User) (*userv1.User, error)
	AddUser(context.Context, *userv1.User) (*userv1.User, error)
	EditUser(context.Context, *userv1.User) (*userv1.User, error)
	RemoveUser(context.Context, *userv1.User) (*userv1.User, error)
}

type gormUserRepo struct {
	db *gorm.DB
}

func NewUserRepository(dbMain *gorm.DB) *gormUserRepo {
	return &gormUserRepo{
		db: dbMain,
	}
}

func (repo *gormUserRepo) GetUserList(ctx context.Context, req *userv1.User, query *commonv1.StandardQuery) ([]*userv1.User, *commonv1.StandardPaginationResponse, error) {
	usersORM := []*userv1.UserORM{}
	if query.Page < 1 {
		query.Page = 1
	}
	pagination := &commonv1.StandardPaginationResponse{
		Page: query.Page,
		Total: 0,
		Found: 0,
	}

	statement := repo.db.Model(&userv1.UserORM{}).
		Where("deleted_at IS NULL").
		Order("created_at DESC")
	if req != nil {
		reqORM, err := req.ToORM(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert request to ORM")
			return nil, pagination, err
		}
		statement = statement.Where(&reqORM)
	}

	if query.Search != "" {
		query.Search = strings.ToLower(query.Search)
		statement = statement.Where("lower(user_name) LIKE ? OR lower(name) LIKE ? OR lower(email) LIKE ?", "%"+query.Search+"%", "%"+query.Search+"%", "%"+query.Search+"%")
	}

	statement.Count(&pagination.Total)
	if query != nil && query.Page > 0 && query.PageSize > 0 {
		statement = statement.Offset(int(query.Page - 1) * int(query.PageSize)).Limit(int(query.PageSize))
	}

	err := statement.Count(&pagination.Found).Find(&usersORM).Error
	if err != nil {
		logger.Errorln("Fail to get user list from DB")
		return nil, pagination, err
	}

	users := make([]*userv1.User, len(usersORM))
	for i, userORM := range usersORM {
		userPB, err := userORM.ToPB(ctx)
		if err != nil {
			logger.Fatalln("Fail to convert response to PB")
			return nil, pagination, err
		}
		users[i] = &userPB
	}

	return users, pagination, nil
}

func (repo *gormUserRepo) GetUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
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
		if err != gorm.ErrRecordNotFound {
			logger.Errorln("Fail to get user from DB")
		}
		return nil, err
	}

	userPB, err := userORM.ToPB(ctx)
	if err != nil {
		logger.Fatalln("Fail to convert response to PB")
		return nil, err
	}
	
	return &userPB, nil
}

func (repo *gormUserRepo) AddUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
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

func (repo *gormUserRepo) EditUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
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

func (repo *gormUserRepo) RemoveUser(ctx context.Context, req *userv1.User) (*userv1.User, error) {
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
