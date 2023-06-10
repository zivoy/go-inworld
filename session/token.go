package session

import (
	"time"
)

const ExpirationDelta = 5 * time.Minute

type Token struct {
	Token          string
	Type           string
	ExpirationTime time.Time
	SessionId      string
}

func (s *Token) IsExpired() bool {
	return s.ExpirationTime.Sub(time.Now()).Abs() <= ExpirationDelta
}

func (s *Token) Equal(o *Token) bool {
	return (s == o) || s != nil && o != nil &&
		s.SessionId == o.SessionId &&
		s.Type == o.Type &&
		s.Token == o.Token &&
		s.ExpirationTime == o.ExpirationTime
}
