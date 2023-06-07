package grpc

import (
	goinworld "github.com/zivoy/go-inworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var UserAgent = "go-inworld/"

func init() {
	UserAgent = UserAgent + "v1.0.0" // todo get the useragent from a build config or something
}

type grpcClient interface {
	Close() error
	Open(config *goinworld.Config) error
	setConn(conn *grpc.ClientConn)
}

func openConnection(client grpcClient, host string, ssl bool) error {
	var creds credentials.TransportCredentials
	if ssl {
		creds = credentials.NewTLS(nil) // i think this is the same as the js credentials.createSsl()
	} else {
		creds = insecure.NewCredentials()
	}

	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(creds), grpc.WithUserAgent(UserAgent))
	client.setConn(conn)
	return err
}
