package core

import (
	"context"

	log "github.com/sirupsen/logrus"

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
			log.Printf(`%v: %v`, req.Peer().Protocol, req.Spec().Procedure)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
