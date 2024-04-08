package main

import (
	"github.com/sandisuryadi36/sansan-dashboard/apps"
	"github.com/sandisuryadi36/sansan-dashboard/core"
	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	"github.com/sandisuryadi36/sansan-dashboard/gen/user/v1/userv1connect"
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
		&handler.UserServiceHandler{
			Repo: *repository.NewUserRepository(core.DBMain),
		},
	)
	path, handler := userv1connect.NewUserServiceHandler(serviceHandler.ServiceHandler.(*handler.UserServiceHandler))

	// run the server
	apps.RunServer(apps.ServerSpec{
		RpcPath:                            path,
		RpcHandler:                         handler,
		RegisterServiceHandlerFromEndpoint: userv1.RegisterUserServiceHandlerFromEndpoint,
		RpcPort:                            rpcPort,
		HttpPort:                           httpPort,
	})

	core.CloseDBMain()
}
