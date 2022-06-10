package main

import (
	"git.jetbrains.space/artdecoction/gt/drun/drun"
	"git.jetbrains.space/artdecoction/wt/tower/lib/config"
	"git.jetbrains.space/artdecoction/wt/tower/lib/fauth"
	"git.jetbrains.space/artdecoction/wt/tower/lib/fauth/claims"
	"git.jetbrains.space/artdecoction/wt/tower/lib/logs"
	"git.jetbrains.space/artdecoction/wt/tower/lib/tower"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/accountsserver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	config.Init()

	app := tower.NewTowerApp()
	defer tower.CleanupApp(app)

	server := newRpcServer(app)
	port := ":" + config.Get().Port

	err := drun.RpcRuntime(server, port)
	if err != nil {
		app.Logger.Error("rpc server error", zap.Error(err))
	}
}

func newRpcServer(app *tower.App) *grpc.Server {
	unaryInterceptors := grpc.ChainUnaryInterceptor(
		logs.UnaryInfoInterceptor(app.Logger),
		fauth.UnaryAuthInterceptor(
			config.Get().AuthenticationMockEnabled,
			claims.BasicClaims{},
			app.FirebaseClients.Auth,
			[]string{},
		),
	)

	opts := []grpc.ServerOption{unaryInterceptors}

	server := grpc.NewServer(opts...)

	accountsserver.RegisterNewAccountsRpcPublicServer(server, app)

	return server
}
