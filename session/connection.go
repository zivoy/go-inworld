package session

import (
	"context"
	"errors"
	goinworld "github.com/zivoy/go-inworld"
	"github.com/zivoy/go-inworld/internal"
	"github.com/zivoy/go-inworld/internal/grpc"
	"github.com/zivoy/go-inworld/internal/protoBuf/packets"
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
	config *goinworld.ClientConfig
	apiKey *goinworld.ApiKey

	token *Token
	scene *Scene

	tokenClient  *grpc.TokenClient
	engineClient *grpc.WorldClient
	stream       *grpc.BilateralWorldSession
	queue        internal.Queue[*packets.InworldPacket] // is this really needed?

	done        chan interface{}
	messageSent chan interface{}

	eventGenerator *grpc.Generator
	OnDisconnect   func()
	OnError        func(err error) // maybe make it so all errors go to this function
	OnMessage      func(packet *goinworld.InworldPacket)
}

func NewConnection() *Connection { //todo
	return &Connection{
		state: connectionInactive,

		done:        make(chan interface{}),
		messageSent: make(chan interface{}),

		eventGenerator: &grpc.Generator{},
	}
}

func (c *Connection) IsActive() bool { return c.state == connectionActive }
func (c *Connection) IsAutoReconnect() bool {
	return c.config != nil &&
		c.config.Connection != nil && //todo check if these are actually not set at any point
		c.config.Connection.AutoReconnect
}

func (c *Connection) GenerateSessionToken(ctx context.Context) (*Token, error) {
	ses, err := c.tokenClient.GenerateSessionToken(ctx, c.apiKey)
	if err != nil {
		return nil, err
	}
	return internal.GetTokenFromProto(ses), nil
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

func (c *Connection) GetCharactersList(ctx context.Context) ([]*Character, error) {
	if c.scene == nil {
		err := c.loadScene(ctx)
		if err != nil {
			return nil, err
		}
	}

	return c.scene.Characters, nil
}

func (c *Connection) Open(ctx context.Context) error {
	err := c.loadScene(ctx)
	if err != nil {
		return err
	}
	if c.state == connectionLoading {
		return nil
	}

	if c.state == connectionLoaded {
		c.state = connectionActivating

		c.stream, err = c.engineClient.Session(ctx, c.token)
		if err != nil {
			c.queue.Clear()
			return err
		}

		c.state = connectionActive
	}

	go c.startListener()
	go c.timeoutLoop()

	// clear queue
	for c.queue.Size() > 0 {
		if err := c.send(c.queue.Pop()); err != nil {
			c.queue.Clear()
			return err // todo maybe throw to the on error function?
		}
	}
	return nil
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
				go c.OnMessage(internal.GetMessageFromProto(message))
			}
		}
	}
}

func (c *Connection) timeoutLoop() {
	if c.config == nil || c.config.Connection == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.config.Connection.DisconnectTimout)
	for {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.Canceled) {
				continue
			}
			go c.Close()
			cancel()
			return
		case <-c.messageSent:
			cancel()
			ctx, cancel = context.WithTimeout(context.Background(), c.config.Connection.DisconnectTimout)
		}
	}
}

func (c *Connection) send(packet *packets.InworldPacket) error {
	if !c.IsActive() && !c.IsAutoReconnect() {
		return errors.New("can't send, not connected")
	}

	if c.stream == nil {
		c.queue.Append(packet)
		return c.Open(context.Background()) // open on background context since a context wasn't proved with open
	}
	c.messageSent <- true // todo maybe put this in goroutine in case it blocks?
	return c.stream.Send(packet)
}

func (c *Connection) loadScene(ctx context.Context) error {
	if c.state == connectionLoading {
		return nil
	}

	// maybe make it so session can be saved -- not needed
	token, err := c.getSessionToken(ctx)
	if err != nil {
		return err
	}
	c.token = token

	if c.scene == nil {

	}

	//todo
}

func (c *Connection) getSessionToken(ctx context.Context) (*Token, error) {
	token := c.token

	// need new token
	if token == nil || token.IsExpired() {
		c.state = connectionLoading

		var err error
		token, err = c.GenerateSessionToken(ctx)
		if err != nil {
			return nil, err
		}
		if c.token != nil {
			token.SessionId = c.token.SessionId
		}
	}

	return token, nil
}

func (c *Connection) getScene(ctx context.Context) (*Scene, error) {

	s, err := c.engineClient.LoadScene(ctx)
	if err != nil {
		return nil, err
	}
	scene := internal.GetSceneFromProto(s)

	return scene, nil
}

////

// SetCurrentCharacter sets which character you are currently talking to
func (c *Connection) SetCurrentCharacter(characterId string) {
	c.eventGenerator.TargetId = characterId
}

func (c *Connection) SendText(text string) error {
	return c.send(c.eventGenerator.NewTextPacket(text))
}
func (c *Connection) SendAudio(chunk []byte) error {
	return c.send(c.eventGenerator.NewDataChunkPacket(chunk, packets.DataChunk_AUDIO))
}

func (c *Connection) SendTrigger(name string, parameters ...goinworld.TriggerParameter) error {
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
