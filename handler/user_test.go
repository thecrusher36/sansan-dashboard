package handler

import (
	"context"
	"errors"
	"testing"

	"connectrpc.com/connect"
	"github.com/sandisuryadi36/sansan-dashboard/core/caller"
	commonv1 "github.com/sandisuryadi36/sansan-dashboard/gen/common/v1"
	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	repo "github.com/sandisuryadi36/sansan-dashboard/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserServiceHandler_GetUser(t *testing.T) {
	cntrl := gomock.NewController(t)
	ctx := context.Background()

	repoControler := repo.NewMockUserRepository(cntrl)
	callerControler := caller.NewMockServiceCaller(cntrl)
	userServiceMock := NewUserHandler(repoControler, callerControler)

	repoControler.EXPECT().GetUser(ctx, &userv1.User{Id: 5}).
		Return(&userv1.User{}, connect.NewError(connect.CodeNotFound, errors.New("user not found")))
	_, err := userServiceMock.GetUser(ctx, connect.NewRequest(&userv1.GetUserRequest{Id: 5}))
	assert.Error(t, err, "user not found")

	repoControler.EXPECT().GetUser(ctx, &userv1.User{Id: 5}).
		Return(&userv1.User{
			Id:   5,
			Name: "test",
		}, nil)
	_, err = userServiceMock.GetUser(ctx, connect.NewRequest(&userv1.GetUserRequest{Id: 5}))
	assert.NoError(t, err)
}

func TestUserServiceHandler_GetUserList(t *testing.T) {
	cntrl := gomock.NewController(t)
	ctx := context.Background()

	repoControler := repo.NewMockUserRepository(cntrl)
	callerControler := caller.NewMockServiceCaller(cntrl)
	userServiceMock := NewUserHandler(repoControler, callerControler)

	repoControler.EXPECT().GetUserList(ctx, &userv1.User{}, &commonv1.StandardQuery{}).
		Return([]*userv1.User{
			{
				Id:   5,
				Name: "test",
			},
			{
				Id:   6,
				Name: "david",
			},
			{
				Id:   8,
				Name: "jonny",
			},
		}, &commonv1.StandardPaginationResponse{Page: 1, Total: 3, Found: 3}, nil)
	res, err := userServiceMock.GetUserList(ctx, connect.NewRequest(&userv1.GetUserListRequest{Query: &commonv1.StandardQuery{}}))
	assert.NoError(t, err)
	assert.Len(t, res.Msg.Users, 3)
}
