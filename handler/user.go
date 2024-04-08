package handler

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	"github.com/sandisuryadi36/sansan-dashboard/gen/user/v1/userv1connect"
	"github.com/sandisuryadi36/sansan-dashboard/repository"
)

type UserHandler interface {
	GetUserList(context.Context, *connect.Request[userv1.GetUserListRequest]) (*connect.Response[userv1.GetUserListResponse], error)
	AddUser(context.Context, *connect.Request[userv1.AddUserRequest]) (*connect.Response[userv1.AddUserResponse], error)
}

type UserServiceHandler struct {
	userv1connect.UnimplementedUserServiceHandler
	Repo repository.GormUserRepo
}
