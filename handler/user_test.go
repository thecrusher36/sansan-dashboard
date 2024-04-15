package handler

import (
	"context"
	"errors"
	"testing"

	"connectrpc.com/connect"
	"github.com/sandisuryadi36/sansan-dashboard/core/caller"
	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	"github.com/sandisuryadi36/sansan-dashboard/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserServiceHandler_GetUser(t *testing.T) {
	cntrl := gomock.NewController(t)
	ctx := context.Background()

	repoControler := mock.NewMockUserRepository(cntrl)
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
