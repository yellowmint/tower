package accounts

import (
	"context"
	"fmt"
	rpcpublicv1 "git.jetbrains.space/artdecoction/wt/tower/contracts/accounts/rpcpublic/v1"
	"git.jetbrains.space/artdecoction/wt/tower/integrationtests/support"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func TestAccounts(t *testing.T) {
	s := support.Init()
	defer s.Cleanup()

	cc := s.NewGrpcClientConn(t, "accounts")
	client := rpcpublicv1.NewAccountsServiceClient(cc)

	authUserId := fmt.Sprintf("%d", time.Now().UTC().Unix())
	name := "Tommy55"
	accountId := ""

	t.Run("GetAccount not found", func(t *testing.T) {
		ctx := s.AuthorizeInContext(context.Background(), authUserId, uuid.New())

		request := &rpcpublicv1.GetAccountRequest{AccountId: "9967a44b-f26c-486f-89cb-d048afa0c38b"}

		_, err := client.GetAccount(ctx, request)
		assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
	})

	t.Run("GetMyAccount not found", func(t *testing.T) {
		ctx := s.AuthorizeInContext(context.Background(), authUserId, uuid.New())

		_, err := client.GetMyAccount(ctx, &rpcpublicv1.GetMyAccountRequest{})
		assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
	})

	t.Run("CreateMyAccount success", func(t *testing.T) {
		ctx := s.AuthorizeInContext(context.Background(), authUserId, uuid.New())

		request := &rpcpublicv1.CreateMyAccountRequest{Name: name}

		_, err := client.CreateMyAccount(ctx, request)
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
	})

	t.Run("GetMyAccount success", func(t *testing.T) {
		ctx := s.AuthorizeInContext(context.Background(), authUserId, uuid.New())

		res, err := client.GetMyAccount(ctx, &rpcpublicv1.GetMyAccountRequest{})
		assert.Equal(t, codes.OK.String(), status.Code(err).String())

		assert.Equal(t, name, res.Name)
		assert.NotEmpty(t, res.AccountId)
		accountId = res.AccountId
	})

	t.Run("GetAccount success", func(t *testing.T) {
		ctx := s.AuthorizeInContext(context.Background(), authUserId, uuid.New())

		request := &rpcpublicv1.GetAccountRequest{AccountId: accountId}

		res, err := client.GetAccount(ctx, request)
		assert.Equal(t, codes.OK.String(), status.Code(err).String())

		assert.Equal(t, name, res.Name)
		assert.Equal(t, accountId, res.AccountId)
	})

	t.Run("DeleteMyAccount success", func(t *testing.T) {
		t.Skip()
		ctx := s.AuthorizeInContext(context.Background(), authUserId, uuid.New())

		_, err := client.DeleteMyAccount(ctx, &rpcpublicv1.DeleteMyAccountRequest{})
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
	})

	t.Run("GetAccount not found deleted account", func(t *testing.T) {
		t.Skip()
		ctx := s.AuthorizeInContext(context.Background(), authUserId, uuid.New())

		request := &rpcpublicv1.GetAccountRequest{AccountId: accountId}

		_, err := client.GetAccount(ctx, request)
		assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
	})
}
