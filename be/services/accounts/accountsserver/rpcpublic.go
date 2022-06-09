package accountsserver

import (
	rpcpublicv1 "git.jetbrains.space/artdecoction/wt/tower/contracts/accounts/rpcpublic/v1"
	"git.jetbrains.space/artdecoction/wt/tower/lib/tower"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/repository"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/transport"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/pkg/account"
	"google.golang.org/grpc"
)

func RegisterNewAccountsRpcPublicServer(server *grpc.Server, app *tower.App) {
	accountRepo := repository.NewFirestoreAccountRepo(app.FirestoreClient)
	accountService := account.NewService(app, accountRepo)
	accountServer := transport.NewRpcPublicServer(accountService)

	rpcpublicv1.RegisterAccountsServiceServer(server, accountServer)
}
