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
	config *ClientConfig
	apiKey *ApiKey

	Session *session.Session
	token   *session.Token
	scene   *session.Scene

	tokenClient  *grpc.TokenClient
	engineClient *grpc.WorldClient
	stream       *grpc.BilateralWorldSession
	queue        internal.Queue[*packets.InworldPacket]

	done        chan interface{}
	messageSent chan interface{}

	eventGenerator *grpc.Generator
	OnDisconnect   func()
	OnError        func(err error)
	OnMessage      func(packet *InworldPacket)
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

func (c *Connection) Close() {
	c.done <- true
	c.state = connectionInactive
	err := c.stream.Stop()
	if err != nil {
		c.OnError(err)
	}
	c.stream = nil
	c.queue.Clear()

	if c.OnDisconnect != nil {
		c.OnDisconnect()
	}
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

		c.state = connectionActive
		// todo send data in queue
		// todo start timeout
	}

	go c.startListener()
}

func (c *Connection) startListener() {
	for {
		select {
		case <-c.done:
			return

		case err := <-c.stream.Done:
			if err != nil {
				go c.OnError(err)
			}
			go c.Close()

		case message := <-c.stream.Incoming:
			if c.OnMessage != nil {
				go c.OnMessage(getMessageFromProto(message))
			}
		}
	}
}

func (c *Connection) send(packet *packets.InworldPacket) error {
	if c.stream == nil {
		c.queue.Append(packet)
		return nil
	}
	return c.stream.Send(packet)
	//todo also write function will be part of send
}

func (c *Connection) loadScene() {
	if c.state == connectionLoading {
		return
	}
	//todo
}

////

func (c *Connection) SetCurrentCharacter(characterId string) {
	c.eventGenerator.TargetId = characterId
}

func (c *Connection) SendText(text string) error {
	return c.send(c.eventGenerator.NewTextPacket(text))
}
func (c *Connection) SendAudio(chunk []byte) error {
	return c.send(c.eventGenerator.NewDataChunkPacket(chunk, packets.DataChunk_AUDIO))
}

func (c *Connection) SendTrigger(name string, parameters ...TriggerParameter) error {
	params := make([]*packets.CustomEvent_Parameter, len(parameters))
	for i, parameter := range parameters {
		params[i] = &packets.CustomEvent_Parameter{
			Name:  parameter.Name,
			Value: parameter.Value,
		}
	}

	return c.send(c.eventGenerator.NewTriggerPacket(name, params))
}

//todo other events
