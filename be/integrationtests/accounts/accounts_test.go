package accounts

import (
	"context"
	rpcpublicv1 "git.jetbrains.space/artdecoction/wt/tower/contracts/accounts/rpcpublic/v1"
	"git.jetbrains.space/artdecoction/wt/tower/integrationtests/support"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestAccounts(t *testing.T) {
	s := support.Init()
	defer s.Cleanup()

	cc := s.NewGrpcClientConn("accounts")
	client := rpcpublicv1.NewAccountsServiceClient(cc)

	name := "Tommy55"

	authUser := s.CreateTestAuthUser(context.Background())
	auth := support.Authorization{
		AuthUserId: authUser.UID,
		AccountId:  uuid.UUID{},
	}

	t.Run("CreateMyAccount success", func(t *testing.T) {
		ctx := s.AuthorizeInContext(context.Background(), auth)

		request := &rpcpublicv1.CreateMyAccountRequest{Name: name}

		_, err := client.CreateMyAccount(ctx, request)
		assert.Equal(t, codes.OK.String(), status.Code(err).String())

		auth.AccountId = s.GetAccountIdByAuthUserId(context.Background(), authUser.UID)
	})

	t.Run("GetMyAccount success", func(t *testing.T) {
		ctx := s.AuthorizeInContext(context.Background(), auth)

		res, err := client.GetMyAccount(ctx, &rpcpublicv1.GetMyAccountRequest{})
		assert.Equal(t, codes.OK.String(), status.Code(err).String())

		assert.Equal(t, name, res.Name.Base)
		assert.NotEmpty(t, res.AccountId)
	})

	t.Run("GetAccount not found", func(t *testing.T) {
		ctx := s.AuthorizeInContext(context.Background(), auth)

		request := &rpcpublicv1.GetAccountRequest{AccountId: "9967a44b-f26c-486f-89cb-d048afa0c38b"}

		_, err := client.GetAccount(ctx, request)
		assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
	})

	t.Run("GetAccount success", func(t *testing.T) {
		ctx := s.AuthorizeInContext(context.Background(), auth)

		request := &rpcpublicv1.GetAccountRequest{AccountId: auth.AccountId.String()}

		res, err := client.GetAccount(ctx, request)
		assert.Equal(t, codes.OK.String(), status.Code(err).String())

		assert.Equal(t, name, res.Name.Base)
		assert.Equal(t, auth.AccountId.String(), res.AccountId)
	})

	t.Run("DeleteMyAccount success", func(t *testing.T) {
		ctx := s.AuthorizeInContext(context.Background(), auth)

		_, err := client.DeleteMyAccount(ctx, &rpcpublicv1.DeleteMyAccountRequest{})
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
	})

	t.Run("GetAccount not found deleted account", func(t *testing.T) {
		auth2 := s.NewFakeAuthorization()
		ctx := s.AuthorizeInContext(context.Background(), auth2)

		request := &rpcpublicv1.GetAccountRequest{AccountId: auth.AccountId.String()}

		_, err := client.GetAccount(ctx, request)
		assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
	})
}
