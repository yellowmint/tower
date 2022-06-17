package support

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var gRpcClientConnections []*grpc.ClientConn

func (s *Support) NewGrpcClientConn(service string) *grpc.ClientConn {
	var opts []grpc.DialOption

	host := viper.GetString(service + ".url")
	opts = append(opts, grpc.WithAuthority(host))

	if viper.GetBool(service + "tls") {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			panic(err)
		}

		cred := credentials.NewTLS(&tls.Config{RootCAs: systemRoots})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	cc, err := grpc.Dial(host, opts...)
	if err != nil {
		panic(err)
	}

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
