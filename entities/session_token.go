package entities

import (
	"github.com/zivoy/go-inworld/internal/protoBuf/studioTokens"
	"time"
)

const ExpirationDelta = 5 * time.Minute

type SessionToken struct {
	Token          string
	Type           string
	ExpirationTime time.Time
	SessionId      string
}

func SessionTokenFromProto(proto *studioTokens.SessionAccessToken) *SessionToken {
	return &SessionToken{
		Token:          proto.GetToken(),
		Type:           proto.GetType(),
		ExpirationTime: proto.GetExpirationTime().AsTime(),
		SessionId:      proto.GetSessionId(),
	}
}

func (s *SessionToken) IsExpired() bool {
	return s.ExpirationTime.Sub(time.Now()).Abs() <= ExpirationDelta
}
