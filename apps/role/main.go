package main

import "github.com/sandisuryadi36/sansan-dashboard/core"

func main() {
	// var rpcPort int = 9090
	// var httpPort int = 8080

	// migrate DB
	core.MigrateDB()

	// start DB connection
	core.StartDBConnection()

	// // initiate RPC path and handler
	// authHandler := handler.NewHandler(
	// 	nil,
	// 	&auth.AuthServiceHandler{
	// 		Repo: repository.NewGormAuthRepo(core.DBMain),
	// 	})
	// path, handler := authv1connect.NewAuthServiceHandler(authHandler.AuthHandler)

	// // run the server
	// apps.RunServer(apps.ServerSpec{
	// 	RpcPath:                            path,
	// 	RpcHandler:                         handler,
	// 	RegisterServiceHandlerFromEndpoint: authv1.RegisterAuthServiceHandlerFromEndpoint,
	// 	RpcPort:                            rpcPort,
	// 	HttpPort:                           httpPort,
	// })

	core.CloseDBMain()
}
