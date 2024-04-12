package core

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sandisuryadi36/sansan-dashboard/libs/logger"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServerSpec struct {
	RpcPath                            string
	RpcHandler                         http.Handler
	RegisterServiceHandlerFromEndpoint func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	RpcPort                            int
	HttpPort                           int
}

func RunServer(s ServerSpec) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Initiate RPC Mux
	rpcMux := http.NewServeMux()
	rpcMux.Handle(s.RpcPath, s.RpcHandler)

	// Initiate RPC-gateway Mux
	httpMux := runtime.NewServeMux(
		runtime.WithErrorHandler(CustomHTTPError),
	)

	// Register HTTP handler for RPC service
	err := s.RegisterServiceHandlerFromEndpoint(ctx, httpMux, fmt.Sprintf("localhost:%d", s.RpcPort), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		logger.Fatalf("Failed to register HTTP gateway: %v", err)
	}

	// Initiate HTTP server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", s.HttpPort),
		Handler: httpMux,
	}

	// Initiate listener for HTTP gateway
	httpListener, err := net.Listen("tcp", fmt.Sprintf(":%v", s.HttpPort))
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}

	// serve RPC
	go func() {
		logger.Printf("Starting RPC server on localhost:%v...", s.RpcPort)
		err := http.ListenAndServe(
			fmt.Sprintf(`:%v`, s.RpcPort),
			h2c.NewHandler(rpcMux, &http2.Server{}),
		)
		if err != nil {
			logger.Fatalln("Fail to serve the server")
		}
	}()

	// serve RPC-gateway
	go func() {
		logger.Printf("Starting HTTP server on localhost:%v...", s.HttpPort)
		err = httpServer.Serve(httpListener)
		if err != nil {
			logger.Fatalf("Failed to serve HTTP server: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	// Block until a signal is received
	<-ch
}
