package main

import (
	"github.com/sandisuryadi36/sansan-dashboard/apps"
	"github.com/sandisuryadi36/sansan-dashboard/core"
	rolev1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	"github.com/sandisuryadi36/sansan-dashboard/gen/role/v1/rolev1connect"
	"github.com/sandisuryadi36/sansan-dashboard/handler"
	"github.com/sandisuryadi36/sansan-dashboard/repository"
)

func main() {
	var rpcPort int = 9090
	var httpPort int = 8080

	// migrate DB
	core.MigrateDB()

	// start DB connection
	core.StartDBConnection()

	// initiate RPC path and handler
	serviceHandler := handler.NewHandler(
		&handler.RoleServiceHandler{
			Repo: *repository.NewRoleRepository(core.DBMain),
		},
	)
	path, handler := rolev1connect.NewRoleServiceHandler(
		serviceHandler.ServiceHandler.(*handler.RoleServiceHandler),
		apps.NewInterceotors(),
	)

	// run the server
	apps.RunServer(apps.ServerSpec{
		RpcPath:                            path,
		RpcHandler:                         handler,
		RegisterServiceHandlerFromEndpoint: rolev1.RegisterRoleServiceHandlerFromEndpoint,
		RpcPort:                            rpcPort,
		HttpPort:                           httpPort,
	})

	core.CloseDBMain()
}
