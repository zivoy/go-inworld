package goinworld

import (
	"github.com/google/uuid"
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
	InworldControlUnknown = InworldControlType(iota)
	InworldControlInteractionEnd
)

type EmotionalBehavior uint8
type EmotionalStrength uint8

const (
	EmotionBehaviorNeutral = EmotionalBehavior(iota)
	EmotionBehaviorDisgust
	EmotionBehaviorContempt
	EmotionBehaviorBelligerence
	EmotionBehaviorDomineering
	EmotionBehaviorCriticism
	EmotionBehaviorAnger
	EmotionBehaviorTension
	EmotionBehaviorTenseHumor
	EmotionBehaviorDefensiveness
	EmotionBehaviorWhining
	EmotionBehaviorSadness
	EmotionBehaviorStonewalling
	EmotionBehaviorInterest
	EmotionBehaviorValidation
	EmotionBehaviorAffection
	EmotionBehaviorHumor
	EmotionBehaviorSurprise
	EmotionBehaviorJoy
)

const (
	EmotionalStrengthUnspecified = EmotionalStrength(iota)
	EmotionalStrengthWeak
	EmotionalStrengthStrong
	EmotionalStrengthNormal
)

type ApiKey struct {
	Key    string
	Secret string
}

type ClientConfig struct {
	Connection   *ConnectionConfig
	Capabilities *engine.CapabilitiesRequest //todo replace with something exposable
}

type ConnectionConfig struct {
	AutoReconnect    bool
	DisconnectTimout time.Duration
}

type GenerateSessionTokenFunc = func() func(token session.Token)

type InworldPacket struct {
	PacketType InworldPacketType

	Date     time.Time
	PacketId *PacketId
	Routing  *Routing

	//events
	Text            *TextEvent
	Audio           *AudioEvent
	Control         *ControlEvent
	Trigger         *TriggerEvent
	Emotions        *EmotionEvent
	Silence         *SilenceEvent
	CancelResponses *CancelResponsesEvent
}

type PacketId struct {
	PacketId      uuid.UUID
	UtteranceId   uuid.UUID
	InteractionId uuid.UUID
}

type Routing struct {
	Source *Actor
	Target *Actor
}

type Actor struct {
	Name        string
	IsPlayer    bool
	IsCharacter bool
}

type TextEvent struct {
	Text  string
	Final bool
}
type AudioEvent struct {
	Chunk                 []byte
	AdditionalPhonemeInfo []*AdditionalPhonemeInfo
}
type ControlEvent struct {
	Type InworldControlType
}

type TriggerEvent struct {
	Name       string
	Parameters []*TriggerParameter
}

type EmotionEvent struct {
	Behavior EmotionalBehavior
	Strength EmotionalStrength
}

type SilenceEvent struct {
	Duration time.Duration // should be used in ms
}

type CancelResponsesEvent struct {
	InteractionId uuid.UUID
	UtteranceId   []uuid.UUID
}

type AdditionalPhonemeInfo struct {
	Phoneme     string
	StartOffset time.Duration // should be used in seconds
}

type TriggerParameter struct {
	Name  string
	Value string
}
