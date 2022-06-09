package transport

import (
	"context"
	rpcpublicv1 "git.jetbrains.space/artdecoction/wt/tower/contracts/accounts/rpcpublic/v1"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/pkg/account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type rpcPublicServer struct {
	accountService account.Service
	rpcpublicv1.UnimplementedAccountsServiceServer
}

func NewRpcPublicServer(*account.Svc) rpcpublicv1.AccountsServiceServer {
	return rpcPublicServer{}
}

func (s rpcPublicServer) GetAccount(context.Context, *rpcpublicv1.GetAccountRequest) (*rpcpublicv1.GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountById not implemented")
}

func (s rpcPublicServer) GetMyAccount(context.Context, *rpcpublicv1.GetMyAccountRequest) (*rpcpublicv1.GetMyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyAccount not implemented")
}

func (s rpcPublicServer) CreateMyAccount(context.Context, *rpcpublicv1.CreateMyAccountRequest) (*rpcpublicv1.CreateMyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMyAccount not implemented")
}

func (s rpcPublicServer) DeleteMyAccount(context.Context, *rpcpublicv1.DeleteMyAccountRequest) (*rpcpublicv1.DeleteMyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMyAccount not implemented")
}
