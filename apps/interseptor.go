package apps

import (
	"context"
	"log"

	"connectrpc.com/connect"
)

func NewInterceotors() connect.HandlerOption {
	return connect.WithInterceptors(
		LogInterceptor(),
	)
}

func LogInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			log.Printf(`%v: %v %v`, req.HTTPMethod(), req.Spec().Procedure, req.Peer().Protocol)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
