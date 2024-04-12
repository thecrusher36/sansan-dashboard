package core

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/sandisuryadi36/sansan-dashboard/libs/logger"
)

func NewInterceotors() connect.HandlerOption {
	return connect.WithInterceptors(
		LogInterceptor(),
	)
}

func LogInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()
			logger.Println("rpc start")

			res, err := next(ctx, req)

			if err != nil {
				logger.Errorln(err)
			}

			logger.Printf(`rpc finish: %v, duration: %v`, req.Spec().Procedure, time.Since(start))
			return res, err
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
