package validation

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RpcMessage interface {
	ValidateAll() error
}

func UnaryValidationInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		req, err := tryValidate(req)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		res, err = tryValidate(res)
		if err != nil {
			logger.Error("response validation failed", zap.Error(err))
			return nil, status.Error(codes.Internal, "response validation failed")
		}

		return res, err
	}
}

func tryValidate(msg interface{}) (interface{}, error) {
	msgValidatable, ok := msg.(RpcMessage)
	if !ok {
		return msg, nil
	}

	err := msgValidatable.ValidateAll()
	if err != nil {
		return msg, err
	}

	return msg, nil
}
