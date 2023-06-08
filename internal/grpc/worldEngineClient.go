package grpc

import (
	"context"
	"errors"
	"fmt"
	goinworld "github.com/zivoy/go-inworld"
	"github.com/zivoy/go-inworld/internal/protoBuf/engine"
	"github.com/zivoy/go-inworld/internal/protoBuf/packets"
	"github.com/zivoy/go-inworld/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
)

const ClientId = "go"

var _ grpcClient = &WorldClient{}

type WorldClient struct {
	config *goinworld.Config
	conn   *grpc.ClientConn
	client engine.WorldEngineClient
}

type BilateralWorldSession struct {
	Incoming chan *packets.InworldPacket
	Done     chan error // will send nil when closed gracefully, sends errors

	session engine.WorldEngine_SessionClient
}

func (c *WorldClient) Open(config *goinworld.Config) error {
	c.config = config
	return openConnection(c, config.EngineHost, config.EngineSSL)
}

func (c *WorldClient) setConn(conn *grpc.ClientConn) {
	c.conn = conn
	c.client = engine.NewWorldEngineClient(c.conn)
}

func (c *WorldClient) Close() error {
	if c.conn == nil {
		return nil
	}
	c.client = nil
	return c.conn.Close()
}

func (c *WorldClient) addMetadata(ctx context.Context, token *session.Token) {
	metadata.AppendToOutgoingContext(ctx,
		"session-id", token.SessionId,
		"authorization", fmt.Sprintf("%s %s", token.Type, token.Token))
}

func (c *WorldClient) LoadScene(ctx context.Context,
	name string, token *session.Token, user *engine.UserRequest, capabilities *engine.CapabilitiesRequest, client *engine.ClientRequest) (*engine.LoadSceneResponse, error) {
	if c.conn == nil {
		return nil, errors.New("not connected")
	}
	c.addMetadata(ctx, token)

	cl := client
	if cl == nil {
		cl = &engine.ClientRequest{Id: ClientId}
	}

	return c.client.LoadScene(ctx, &engine.LoadSceneRequest{
		Name:         name,
		Capabilities: capabilities,
		User:         user,
		Client:       cl,
	})
}

func (c *WorldClient) Scene(ctx context.Context) (*BilateralWorldSession, error) {
	session, err := c.client.Session(ctx)
	if err != nil {
		return nil, err
	}

	ses := &BilateralWorldSession{
		Incoming: make(chan *packets.InworldPacket),
		Done:     make(chan error),
		session:  session,
	}

	go ses.startListening()

	return ses, nil
}

func (s *BilateralWorldSession) startListening() {
	for {
		req, err := s.session.Recv()
		if err == io.EOF {
			s.Done <- nil
			return
		}
		if err != nil {
			log.Printf("receive error %v", err)
			s.Done <- errors.New("failed to receive request")
			return
		}

		s.Incoming <- req
	}
}

func (s *BilateralWorldSession) Send(packet *packets.InworldPacket) error {
	return s.session.Send(packet)
}

func (s *BilateralWorldSession) Stop() error {
	return s.session.CloseSend()
}
