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

func TestAccountsInterceptors(t *testing.T) {
	s := support.Init()
	defer s.Cleanup()

	cc := s.NewGrpcClientConn(t, "accounts")
	client := rpcpublicv1.NewAccountsServiceClient(cc)

	t.Run("authorization", func(t *testing.T) {
		ctx := context.Background()

		request := &rpcpublicv1.GetAccountRequest{}

		_, err := client.GetAccount(ctx, request)
		assert.Equal(t, codes.Unauthenticated.String(), status.Code(err).String())
	})

	t.Run("validation", func(t *testing.T) {
		authUserId := fmt.Sprintf("%d", time.Now().UTC().Unix())
		ctx := s.AuthorizeInContext(context.Background(), authUserId, uuid.New())

		request := &rpcpublicv1.GetAccountRequest{AccountId: "650d957b-8add-9a768129ca49"}

		_, err := client.GetAccount(ctx, request)
		assert.Equal(t, codes.InvalidArgument.String(), status.Code(err).String())
		assert.Contains(t, err.Error(), "desc = invalid GetAccountRequest.AccountId: value must be a valid UUID | caused by: invalid uuid format")
	})
}
