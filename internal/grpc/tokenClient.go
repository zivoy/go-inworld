package grpc

import (
	"context"
	"errors"
	goinworld "github.com/zivoy/go-inworld"
	"github.com/zivoy/go-inworld/internal/auth"
	"github.com/zivoy/go-inworld/protoBuf/studioTokens"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var _ grpcClient = &TokenClient{}

type TokenClient struct {
	config *goinworld.Config
	conn   *grpc.ClientConn
	client studioTokens.TokensClient
}

func (c *TokenClient) Open(config *goinworld.Config) error {
	c.config = config
	return openConnection(c, config.StudioHost, config.StudioSSL)
}

func (c *TokenClient) setConn(conn *grpc.ClientConn) {
	c.conn = conn
	c.client = studioTokens.NewTokensClient(c.conn)
}

func (c *TokenClient) Close() error {
	if c.conn == nil {
		return nil
	}
	c.client = nil
	return c.conn.Close()
}

func (c *TokenClient) GenerateSessionToken(ctx context.Context, key goinworld.ApiKey) (*studioTokens.SessionAccessToken, error) {
	if c.conn == nil {
		return nil, errors.New("not connected")
	}

	signature, err := auth.GetAuthorization(c.config.StudioHost, key)
	if err != nil {
		return nil, err
	}

	metadata.AppendToOutgoingContext(ctx, "authorization", signature)
	return c.client.GenerateSessionToken(ctx, &studioTokens.GenerateSessionTokenRequest{Key: key.Key})
}
