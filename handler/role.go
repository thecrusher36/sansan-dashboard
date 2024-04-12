package handler

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	commonv1 "github.com/sandisuryadi36/sansan-dashboard/gen/common/v1"
	rolev1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	"github.com/sandisuryadi36/sansan-dashboard/gen/role/v1/rolev1connect"
	"github.com/sandisuryadi36/sansan-dashboard/repository"
)

type RoleHandler interface {
	GetRoleList(context.Context, *connect.Request[rolev1.GetRoleListRequest]) (*connect.Response[rolev1.GetRoleListResponse], error)
}

type RoleServiceHandler struct {
	rolev1connect.UnimplementedRoleServiceHandler
	Repo repository.GormRoleRepo
}

func (h *RoleServiceHandler) GetRoleList(ctx context.Context, req *connect.Request[rolev1.GetRoleListRequest]) (res *connect.Response[rolev1.GetRoleListResponse], err error) {
	roles, err := h.Repo.GetRoleList(ctx)
	if err != nil {
		return
	}

	res = connect.NewResponse(&rolev1.GetRoleListResponse{
		Roles: roles,
		HttpStatus: &commonv1.StandardResponse{
			Message: "success",
			Code:    http.StatusOK,
		},
	})

	return
}
