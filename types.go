package goinworld

import (
	"github.com/zivoy/go-inworld/entities"
	"github.com/zivoy/go-inworld/protoBuf/engine"
	"time"
)

type ApiKey struct {
	Key    string
	Secret string
}

type ClientConfiguration struct {
	Connection   *ConnectionConfig
	Capabilities engine.CapabilitiesRequest
}

type ConnectionConfig struct {
	AutoReconnect    bool
	disconnectTimout time.Duration
}

type GenerateSessionTokenFunc = func() func(token entities.SessionToken)
