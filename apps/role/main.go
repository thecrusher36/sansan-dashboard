package main

import (
	"github.com/sandisuryadi36/sansan-dashboard/core"
	rolev1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	"github.com/sandisuryadi36/sansan-dashboard/gen/role/v1/rolev1connect"
	"github.com/sandisuryadi36/sansan-dashboard/handler"
	"github.com/sandisuryadi36/sansan-dashboard/repository"
	"gorm.io/gorm/logger"
)

func main() {
	var rpcPort int = 9090
	var httpPort int = 8080

	// migrate DB
	core.MigrateDB()

	// start DB connection
	core.StartDBConnection()
	core.DBMain.Config.Logger.LogMode(logger.Info)
	defer core.CloseDBMain()

	// initiate RPC path and handler
	serviceHandler := handler.NewRoleHandler(handler.RoleServiceHandler{
		Repo: repository.NewRoleRepository(core.DBMain),
	})
	path, handler := rolev1connect.NewRoleServiceHandler(
		serviceHandler,
		core.NewInterceotors(),
	)

	// run the server
	core.RunServer(core.ServerSpec{
		RpcPath:                            path,
		RpcHandler:                         handler,
		RegisterServiceHandlerFromEndpoint: rolev1.RegisterRoleServiceHandlerFromEndpoint,
		RpcPort:                            rpcPort,
		HttpPort:                           httpPort,
	})

}
