package accounts

import (
	"crypto/tls"
	"crypto/x509"
	rpcpublicv1 "git.jetbrains.space/artdecoction/wt/tower/contracts/accounts/rpcpublic/v1"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func newAccountsClient(t *testing.T) (*grpc.ClientConn, rpcpublicv1.AccountsServiceClient) {
	var opts []grpc.DialOption

	host := viper.GetString("accounts.url")
	opts = append(opts, grpc.WithAuthority(host))

	if viper.GetBool("accounts.tls") {
		systemRoots, err := x509.SystemCertPool()
		assert.NoError(t, err)

		cred := credentials.NewTLS(&tls.Config{RootCAs: systemRoots})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	cc, err := grpc.Dial(host, opts...)
	assert.NoError(t, err)

	client := rpcpublicv1.NewAccountsServiceClient(cc)

	return cc, client
}

func closeClient(t *testing.T, cc *grpc.ClientConn) {
	err := cc.Close()
	assert.NoError(t, err)
}
