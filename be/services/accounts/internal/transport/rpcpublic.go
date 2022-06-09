package transport

import (
	"context"
	rpcpublicv1 "git.jetbrains.space/artdecoction/wt/tower/contracts/accounts/rpcpublic/v1"
	"git.jetbrains.space/artdecoction/wt/tower/lib/fauth/claims"
	"git.jetbrains.space/artdecoction/wt/tower/lib/tower"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/pkg/account"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type rpcPublicServer struct {
	app            *tower.App
	accountService account.Service
	rpcpublicv1.UnimplementedAccountsServiceServer
}

func NewRpcPublicServer(app *tower.App, service *account.Svc) rpcpublicv1.AccountsServiceServer {
	return &rpcPublicServer{
		app:            app,
		accountService: service,
	}
}

func (s rpcPublicServer) GetAccount(ctx context.Context, req *rpcpublicv1.GetAccountRequest) (*rpcpublicv1.GetAccountResponse, error) {
	accountId := uuid.MustParse(req.GetAccountId())

	res, err := s.accountService.Get(ctx, accountId)
	if err != nil {
		return nil, s.translateError(err)
	}

	return &rpcpublicv1.GetAccountResponse{
		AccountId: res.AccountId.String(),
		Name:      res.Name,
	}, nil
}

func (s rpcPublicServer) GetMyAccount(ctx context.Context, req *rpcpublicv1.GetMyAccountRequest) (*rpcpublicv1.GetMyAccountResponse, error) {
	userId := claims.GetUserIdFromCtx(ctx)

	res, err := s.accountService.GetByUserId(ctx, uuid.MustParse(userId))
	if err != nil {
		return nil, s.translateError(err)
	}

	return &rpcpublicv1.GetMyAccountResponse{
		AccountId: res.AccountId.String(),
		Name:      res.Name,
	}, nil
}

func (s rpcPublicServer) CreateMyAccount(ctx context.Context, req *rpcpublicv1.CreateMyAccountRequest) (*rpcpublicv1.CreateMyAccountResponse, error) {
	userId := claims.GetUserIdFromCtx(ctx)

	err := s.accountService.Create(ctx, uuid.MustParse(userId), req.GetName())
	if err != nil {
		return nil, s.translateError(err)
	}

	return &rpcpublicv1.CreateMyAccountResponse{}, nil
}

func (s rpcPublicServer) DeleteMyAccount(ctx context.Context, req *rpcpublicv1.DeleteMyAccountRequest) (*rpcpublicv1.DeleteMyAccountResponse, error) {
	userId := claims.GetUserIdFromCtx(ctx)

	err := s.accountService.DeleteByUserId(ctx, uuid.MustParse(userId))
	if err != nil {
		return nil, s.translateError(err)
	}

	return &rpcpublicv1.DeleteMyAccountResponse{}, nil
}

func (s rpcPublicServer) translateError(err error) error {
	towerError, ok := err.(tower.Err)
	if !ok {
		s.app.Logger.Error("internal server error", zap.Error(err))
		return status.Error(codes.Internal, "internal server error")
	}

	if towerError == account.ErrAccountNotFound {
		return status.Error(codes.NotFound, towerError.EndUserMessage)
	}
	if towerError == account.ErrAccountAlreadyCreated {
		return status.Error(codes.AlreadyExists, towerError.EndUserMessage)
	}

	s.app.Logger.Error("internal server error", zap.Error(err))
	return status.Error(codes.Internal, "internal server error")
}