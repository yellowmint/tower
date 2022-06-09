package main

import (
	"git.jetbrains.space/artdecoction/gt/drun/drun"
	"git.jetbrains.space/artdecoction/wt/tower/lib/config"
	"git.jetbrains.space/artdecoction/wt/tower/lib/logs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	config.Init()

	logger := logs.NewLogger(config.Get().LoggerFormat)
	defer logs.SyncLogger(logger)

	printStartInfo(logger)

	server := newRpcServer()

	err := drun.RpcRuntime(server, ":"+config.Get().Port)
	if err != nil {
		logger.Error("rpc server error", zap.Error(err))
	}
}

func newRpcServer() *grpc.Server {
	unaryInterceptors := grpc.ChainUnaryInterceptor()

	opts := []grpc.ServerOption{unaryInterceptors}

	server := grpc.NewServer(opts...)

	return server
}

func printStartInfo(logger *zap.Logger) {
	logger.Info("Tower server starting", zap.String("version", config.Get().Version))
}
