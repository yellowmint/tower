package accounts

import (
	"context"
	rpcpublicv1 "git.jetbrains.space/artdecoction/wt/tower/contracts/accounts/rpcpublic/v1"
	"git.jetbrains.space/artdecoction/wt/tower/integrationtests/support"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestAccountsInterceptors(t *testing.T) {
	s := support.Init()
	defer s.Cleanup()

	cc := s.NewGrpcClientConn(t, "accounts")
	client := rpcpublicv1.NewAccountsServiceClient(cc)

	t.Run("GetAccount missing auth", func(t *testing.T) {
		ctx := context.Background()

		request := &rpcpublicv1.GetAccountRequest{}

		_, err := client.GetAccount(ctx, request)
		assert.Equal(t, codes.Unauthenticated.String(), status.Code(err).String())
	})
}
