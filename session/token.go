package session

import (
	"github.com/zivoy/go-inworld/internal/protoBuf/studioTokens"
	"time"
)

const ExpirationDelta = 5 * time.Minute

type Token struct {
	Token          string
	Type           string
	ExpirationTime time.Time
	SessionId      string
}

func TokenFromProto(proto *studioTokens.SessionAccessToken) *Token {
	return &Token{
		Token:          proto.GetToken(),
		Type:           proto.GetType(),
		ExpirationTime: proto.GetExpirationTime().AsTime(),
		SessionId:      proto.GetSessionId(),
	}
}

func (s *Token) IsExpired() bool {
	return s.ExpirationTime.Sub(time.Now()).Abs() <= ExpirationDelta
}
