package api

import (
	goinworld "github.com/zivoy/go-inworld"
	"github.com/zivoy/go-inworld/internal/protoBuf/packets"
	"github.com/zivoy/go-inworld/session"
)

type Session struct {
	client      *goinworld.InworldClient
	connection  *session.Connection
	characterId string
	sceneId     string
	sessionId   string
	serverId    string
	uid         string

	OnDisconnect func()
	OnError      func(err error)
	OnMessage    func(packet packets.InworldPacket) // todo maybe move the packets type out of internal
}

func NewSession(config *goinworld.ClientConfig, apiKey *goinworld.ApiKey,
	uid, sceneId, characterId, playerName, serverId string) *Session {

}

func (s *Session) GetUid() string         { return s.uid }
func (s *Session) GetCharacterId() string { return s.characterId }
func (s *Session) GetSceneId() string     { return s.sceneId }
func (s *Session) GetServerId() string    { return s.serverId }
func (s *Session) GetSessionId() string   { return s.sessionId }

func (s *Session) GetCharacter() *session.Character { //todo
	//if this
	//s.connection
}
func (s *Session) GetCharacters() []*session.Character { //todo
}
