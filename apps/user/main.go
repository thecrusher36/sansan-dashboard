package main

import (
	"github.com/sandisuryadi36/sansan-dashboard/core"
	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	"github.com/sandisuryadi36/sansan-dashboard/gen/user/v1/userv1connect"
	"github.com/sandisuryadi36/sansan-dashboard/handler"
	"github.com/sandisuryadi36/sansan-dashboard/repository"
)

func main() {
	var rpcPort int = 9091
	var httpPort int = 8081

	// migrate DB
	core.MigrateDB()

	// start DB connection
	core.StartDBConnection()
	defer core.CloseDBMain()

	// initiate RPC path and handler
	serviceHandler := handler.NewUserHandler(repository.NewUserRepository(core.DBMain))
	path, handler := userv1connect.NewUserServiceHandler(
		serviceHandler,
		core.NewInterceotors(),
	)

	// run the server
	core.RunServer(core.ServerSpec{
		RpcPath:                            path,
		RpcHandler:                         handler,
		RegisterServiceHandlerFromEndpoint: userv1.RegisterUserServiceHandlerFromEndpoint,
		RpcPort:                            rpcPort,
		HttpPort:                           httpPort,
	})

}
