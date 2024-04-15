package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	commonv1 "github.com/sandisuryadi36/sansan-dashboard/gen/common/v1"
	rolev1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	"github.com/sandisuryadi36/sansan-dashboard/gen/role/v1/rolev1connect"
	"github.com/sandisuryadi36/sansan-dashboard/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RoleHandler interface {
	GetRoleList(context.Context, *connect.Request[rolev1.GetRoleListRequest]) (*connect.Response[rolev1.GetRoleListResponse], error)
	GetRole(context.Context, *connect.Request[rolev1.GetRoleRequest]) (*connect.Response[rolev1.GetRoleResponse], error)
	AddRole(context.Context, *connect.Request[rolev1.AddRoleRequest]) (*connect.Response[rolev1.AddRoleResponse], error)
	EditRole(context.Context, *connect.Request[rolev1.EditRoleRequest]) (*connect.Response[rolev1.EditRoleResponse], error)
	RemoveRole(context.Context, *connect.Request[rolev1.RemoveRoleRequest]) (*connect.Response[rolev1.RemoveRoleResponse], error)
}

type RoleServiceHandler struct {
	rolev1connect.UnimplementedRoleServiceHandler
	Repo repository.RoleRepository
}

func NewRoleHandler(s RoleServiceHandler) *RoleServiceHandler {
	return &RoleServiceHandler{
		Repo: s.Repo,
	}
}

func (h *RoleServiceHandler) GetRoleList(ctx context.Context, req *connect.Request[rolev1.GetRoleListRequest]) (res *connect.Response[rolev1.GetRoleListResponse], err error) {
	roles, err := h.Repo.GetRoleList(ctx, &rolev1.Role{})
	if err != nil {
		return
	}

	res = connect.NewResponse(&rolev1.GetRoleListResponse{
		Roles: roles,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *RoleServiceHandler) GetRole(ctx context.Context, req *connect.Request[rolev1.GetRoleRequest]) (res *connect.Response[rolev1.GetRoleResponse], err error) {
	payload := req.Msg
	role, err := h.Repo.GetRole(ctx, &rolev1.Role{
		Id: payload.Id,
	})
	if err != nil {
		return
	}

	res = connect.NewResponse(&rolev1.GetRoleResponse{
		Role: role,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *RoleServiceHandler) AddRole(ctx context.Context, req *connect.Request[rolev1.AddRoleRequest]) (res *connect.Response[rolev1.AddRoleResponse], err error) {
	payload := req.Msg
	role, err := h.Repo.AddRole(ctx, &rolev1.Role{
		RoleName:        payload.Name,
		RoleDescription: payload.Description,
		CreatedAt:       timestamppb.Now(),
	})
	if err != nil {
		return
	}

	res = connect.NewResponse(&rolev1.AddRoleResponse{
		Role: role,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *RoleServiceHandler) EditRole(ctx context.Context, req *connect.Request[rolev1.EditRoleRequest]) (res *connect.Response[rolev1.EditRoleResponse], err error) {
	payload := req.Msg

	role, err := h.Repo.GetRole(ctx, &rolev1.Role{Id: payload.Id})
	if err != nil {
		return nil, connect.NewError(
			connect.CodeNotFound,
			errors.New("role does not exist"),
		)
	}

	role, err = h.Repo.EditRole(ctx, &rolev1.Role{
		Id:              payload.Id,
		RoleName:        payload.Name,
		RoleDescription: payload.Description,
		CreatedAt:       role.CreatedAt,
	})
	if err != nil {
		return
	}

	res = connect.NewResponse(&rolev1.EditRoleResponse{
		Role: role,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *RoleServiceHandler) RemoveRole(ctx context.Context, req *connect.Request[rolev1.RemoveRoleRequest]) (res *connect.Response[rolev1.RemoveRoleResponse], err error) {
	payload := req.Msg

	role, err := h.Repo.GetRole(ctx, &rolev1.Role{Id: payload.Id})
	if err != nil {
		return nil, connect.NewError(
			connect.CodeNotFound,
			errors.New("role does not exist"),
		)
	}

	_, err = h.Repo.RemoveRole(ctx, &rolev1.Role{Id: role.Id})
	if err != nil {
		return
	}

	res = connect.NewResponse(&rolev1.RemoveRoleResponse{
		Message: fmt.Sprintf(`Role with ID:%v has been removed`, payload.Id),
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}
