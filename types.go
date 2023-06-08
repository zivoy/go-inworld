package goinworld

import (
	"github.com/zivoy/go-inworld/internal/protoBuf/engine"
	"github.com/zivoy/go-inworld/session"
	"time"
)

type ApiKey struct {
	Key    string
	Secret string
}

type ClientConfiguration struct {
	Connection   *ConnectionConfig
	Capabilities *engine.CapabilitiesRequest
}

type ConnectionConfig struct {
	AutoReconnect    bool
	DisconnectTimout time.Duration
}

type GenerateSessionTokenFunc = func() func(token session.Token)
