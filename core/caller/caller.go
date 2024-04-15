package caller

import (
	"net/http"

	"github.com/sandisuryadi36/sansan-dashboard/gen/role/v1/rolev1connect"
	"github.com/sandisuryadi36/sansan-dashboard/gen/user/v1/userv1connect"
)

type ServiceCaller interface {
	Role() rolev1connect.RoleServiceClient
	User() userv1connect.UserServiceClient
}

type caller struct {}

func New() ServiceCaller {
	return &caller{}
}

func (c *caller) Role() rolev1connect.RoleServiceClient {
	return rolev1connect.NewRoleServiceClient(
		http.DefaultClient,
		"http://localhost:9090",
	)
}

func (c *caller) User() userv1connect.UserServiceClient {
	return userv1connect.NewUserServiceClient(
		http.DefaultClient,
		"http://localhost:9091",
	)
}