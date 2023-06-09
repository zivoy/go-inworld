package goinworld

import (
	"context"
	"errors"
	"github.com/zivoy/go-inworld/internal"
	"github.com/zivoy/go-inworld/internal/grpc"
	"github.com/zivoy/go-inworld/internal/protoBuf/packets"
	"github.com/zivoy/go-inworld/session"
)

type connectionState uint8

const (
	connectionInactive = connectionState(iota)
	connectionActive
	connectionActivating
	connectionLoaded
	connectionLoading
)

type Connection struct {
	state  connectionState
	config *ClientConfiguration
	apiKey *ApiKey

	Session *session.Session
	token   *session.Token
	scene   *session.Scene

	tokenClient  *grpc.TokenClient
	engineClient *grpc.WorldClient
	stream       *grpc.BilateralWorldSession
	queue        internal.Queue[packets.InworldPacket]

	EventGenerator *grpc.Generator
}

func NewConnection() *Connection { //todo
	return &Connection{
		state: connectionInactive,
	}
}

func (c *Connection) IsActive() bool { return c.state == connectionActive }
func (c *Connection) IsAutoReconnect() bool {
	return c.config != nil &&
		c.config.Connection != nil && //todo check if these are actually not set at any point
		c.config.Connection.AutoReconnect
}

func (c *Connection) GenerateSessionToken(ctx context.Context) (*session.Token, error) {
	ses, err := c.tokenClient.GenerateSessionToken(ctx, c.apiKey)
	if err != nil {
		return nil, err
	}
	return session.TokenFromProto(ses), nil
}

func (c *Connection) openManually(ctx context.Context) error {
	if c.IsAutoReconnect() {
		return errors.New("can't manually open with autoReconnect on")
	}
	if c.IsActive() {
		return errors.New("already open")
	}

	return c.Open(ctx)
}

func (c *Connection) Close() error {
	c.cancelTimeout()
	c.state = connectionInactive
	//todo
}

func (c *Connection) GetCharactersList() []*session.Character {
	if c.scene == nil {
		c.loadScene()
	}

	return c.scene.Characters
}

func (c *Connection) Open(ctx context.Context) error {
	c.loadScene()

	if c.state == connectionLoaded {
		c.state = connectionActivating

		var err error
		c.stream, err = c.engineClient.Session(ctx, c.token)
		if err != nil {
			return err
		}
		// todo handle incoming data somehow

		c.state = connectionInactive
		// todo send data in queue
		// todo start timeout
	}

}

func (c *Connection) Send(packet *packets.InworldPacket) error {
	//todo also write function will be part of send
}

func (c *Connection) loadScene() {
	if c.state == connectionLoading {
		return
	}
	//todo
}
