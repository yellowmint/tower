package support

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

var gRpcClientConnections []*grpc.ClientConn

func (s *Support) NewGrpcClientConn(t *testing.T, service string) *grpc.ClientConn {
	var opts []grpc.DialOption

	host := viper.GetString(service + ".url")
	opts = append(opts, grpc.WithAuthority(host))

	if viper.GetBool(service + "tls") {
		systemRoots, err := x509.SystemCertPool()
		assert.NoError(t, err)

		cred := credentials.NewTLS(&tls.Config{RootCAs: systemRoots})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	cc, err := grpc.Dial(host, opts...)
	assert.NoError(t, err)

	gRpcClientConnections = append(gRpcClientConnections, cc)

	return cc
}

func closeGrpcClientConnections() {
	for _, connection := range gRpcClientConnections {
		err := connection.Close()
		if err != nil {
			fmt.Printf("error closing grpc client connection %e\n", err)
		}
	}
}
