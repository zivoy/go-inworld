package goinworld

import (
	"github.com/zivoy/go-inworld/internal/protoBuf/engine"
	"github.com/zivoy/go-inworld/session"
	"time"
)

type InworldPacketType uint8
type InworldControlType uint8

const (
	InworldPacketUnknown = InworldPacketType(iota)
	InworldPacketText
	InworldPacketAudio
	InworldPacketTrigger
	InworldPacketEmotion
	InworldPacketControl
	InworldPacketSilence
	InworldPacketCancelResponse
)

const (
	Unknown = InworldControlType(iota)
	InteractionEnd
)

type ApiKey struct {
	Key    string
	Secret string
}

type ClientConfiguration struct {
	Connection   *ConnectionConfig
	Capabilities *engine.CapabilitiesRequest //todo replace with something exposable
}

type ConnectionConfig struct {
	AutoReconnect    bool
	DisconnectTimout time.Duration
}

type GenerateSessionTokenFunc = func() func(token session.Token)

type InworldPacket struct {
	packetType InworldPacketType
}
