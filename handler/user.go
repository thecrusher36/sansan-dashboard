package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/sandisuryadi36/sansan-dashboard/core/caller"
	commonv1 "github.com/sandisuryadi36/sansan-dashboard/gen/common/v1"
	rolev1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	"github.com/sandisuryadi36/sansan-dashboard/gen/user/v1/userv1connect"
	"github.com/sandisuryadi36/sansan-dashboard/libs/auth"
	"github.com/sandisuryadi36/sansan-dashboard/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler interface {
	GetUserList(context.Context, *connect.Request[userv1.GetUserListRequest]) (*connect.Response[userv1.GetUserListResponse], error)
	GetUser(context.Context, *connect.Request[userv1.GetUserRequest]) (*connect.Response[userv1.GetUserResponse], error)
	AddUser(context.Context, *connect.Request[userv1.AddUserRequest]) (*connect.Response[userv1.AddUserResponse], error)
	EditUser(context.Context, *connect.Request[userv1.EditUserRequest]) (*connect.Response[userv1.EditUserResponse], error)
	RemoveUser(context.Context, *connect.Request[userv1.RemoveUserRequest]) (*connect.Response[userv1.RemoveUserResponse], error)
}

type UserServiceHandler struct {
	userv1connect.UnimplementedUserServiceHandler
	Repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserServiceHandler {
	return &UserServiceHandler{
		Repo: repo,
	}
}

func (h *UserServiceHandler) GetUserList(ctx context.Context, req *connect.Request[userv1.GetUserListRequest]) (res *connect.Response[userv1.GetUserListResponse], err error) {
	users, err := h.Repo.GetUserList(ctx, &userv1.User{})
	if err != nil {
		return
	}

	res = connect.NewResponse(&userv1.GetUserListResponse{
		Users: users,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *UserServiceHandler) GetUser(ctx context.Context, req *connect.Request[userv1.GetUserRequest]) (res *connect.Response[userv1.GetUserResponse], err error) {
	payload := req.Msg
	user, err := h.Repo.GetUser(ctx, &userv1.User{
		Id: payload.Id,
	})
	if err != nil {
		return
	}

	res = connect.NewResponse(&userv1.GetUserResponse{
		User: user,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *UserServiceHandler) AddUser(ctx context.Context, req *connect.Request[userv1.AddUserRequest]) (res *connect.Response[userv1.AddUserResponse], err error) {
	payload := req.Msg
	hashedPass, err := auth.HashPassword(payload.Password)
	if err != nil {
		return
	}

	// check if userName or email allready exist
	if err = CheckUserNameAndEmail(ctx, h, &userv1.User{UserName: payload.UserName, Email: payload.Email}); err != nil {
		return
	}

	// check role
	role, err := caller.RoleClient.GetRole(ctx, &connect.Request[rolev1.GetRoleRequest]{
		Msg: &rolev1.GetRoleRequest{
			Id: payload.RoleId,
		},
	})
	if err != nil {
		return nil, connect.NewError(
			connect.CodeNotFound,
			errors.New("role does not exist"),
		)
	}

	user, err := h.Repo.AddUser(ctx, &userv1.User{
		UserName:       payload.UserName,
		Email:          payload.Email,
		Name:           payload.Name,
		HashedPassword: hashedPass,
		Role:           role.Msg.Role,
		CreatedAt:      timestamppb.Now(),
	})
	if err != nil {
		return
	}

	res = connect.NewResponse(&userv1.AddUserResponse{
		User: user,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *UserServiceHandler) EditUser(ctx context.Context, req *connect.Request[userv1.EditUserRequest]) (res *connect.Response[userv1.EditUserResponse], err error) {
	payload := req.Msg
	user, err := h.Repo.GetUser(ctx, &userv1.User{
		Id: payload.Id,
	})
	if err != nil {
		return
	}

	// check if userName or email allready exist
	if err = CheckUserNameAndEmail(ctx, h, &userv1.User{UserName: payload.UserName, Email: payload.Email}); err != nil {
		return
	}

	// check role
	role, err := caller.RoleClient.GetRole(ctx, &connect.Request[rolev1.GetRoleRequest]{
		Msg: &rolev1.GetRoleRequest{
			Id: payload.RoleId,
		},
	})
	if err != nil {
		return nil, connect.NewError(
			connect.CodeNotFound,
			errors.New("role does not exist"),
		)
	}

	user, err = h.Repo.EditUser(ctx, &userv1.User{
		Id:             payload.Id,
		UserName:       payload.UserName,
		Email:          payload.Email,
		Name:           payload.Name,
		HashedPassword: user.HashedPassword,
		Role:           role.Msg.Role,
		CreatedAt:      user.CreatedAt,
	})
	if err != nil {
		return
	}

	res = connect.NewResponse(&userv1.EditUserResponse{
		User: user,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *UserServiceHandler) RemoveUser(ctx context.Context, req *connect.Request[userv1.RemoveUserRequest]) (res *connect.Response[userv1.RemoveUserResponse], err error) {
	payload := req.Msg
	user, err := h.Repo.GetUser(ctx, &userv1.User{
		Id: payload.Id,
	})
	if err != nil {
		return
	}

	_, err = h.Repo.RemoveUser(ctx, &userv1.User{
		Id: user.Id,
	})
	if err != nil {
		return
	}

	res = connect.NewResponse(&userv1.RemoveUserResponse{
		Message: fmt.Sprintf(`User with ID:%v has been removed`, payload.Id),
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

// function to check userName or email allready exist
func CheckUserNameAndEmail(ctx context.Context, h *UserServiceHandler, user *userv1.User) error {
	// check if userName already exist
	if _, err := h.Repo.GetUser(ctx, &userv1.User{
		UserName: user.UserName,
	}); err == nil {
		// logger.Errorln("username already exist")
		return connect.NewError(connect.CodeAlreadyExists, errors.New("username already exist"))
	}

	// check if email already exist
	if _, err := h.Repo.GetUser(ctx, &userv1.User{
		Email: user.Email,
	}); err == nil {
		// logger.Errorln("email already exist")
		return connect.NewError(connect.CodeAlreadyExists, errors.New("email already exist"))
	}

	return nil
}
