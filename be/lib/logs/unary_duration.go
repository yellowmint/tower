package logs

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func UnaryInfoInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now().UTC()

		res, err := handler(ctx, req)

		duration := time.Now().UTC().Sub(startTime)
		logger.Debug(
			"unary call "+info.FullMethod,
			zap.String("took", duration.String()),
			zap.Error(err),
		)

		return res, err
	}
}
