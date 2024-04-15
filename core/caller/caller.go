package caller

import (
	"net/http"

	"github.com/sandisuryadi36/sansan-dashboard/gen/role/v1/rolev1connect"
	"github.com/sandisuryadi36/sansan-dashboard/gen/user/v1/userv1connect"
)

var (
	RoleClient = rolev1connect.NewRoleServiceClient(
		http.DefaultClient,
		"http://localhost:9090",
	)

	UserClient = userv1connect.NewUserServiceClient(
		http.DefaultClient,
		"http://localhost:9091",
	)
)
