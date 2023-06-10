package main

import (
	"errors"
	goinworld "github.com/zivoy/go-inworld"
	"github.com/zivoy/go-inworld/api"
	"github.com/zivoy/go-inworld/internal/protoBuf/engine"
	"log"
	"sync"
)

var app *Application

type Application struct {
	Sessions  map[string]*api.Session
	sessionMu sync.RWMutex
}

func GetApp() *Application {
	if app != nil {
		return app
	}

	app = &Application{
		Sessions: make(map[string]*api.Session, 0),
	}
	return app
}

func (a *Application) OpenSession(request *OpenSessionRequest) (*OpenSessionResponse, error) {
	if _, ok := a.findSession(request.Uid, request.CharacterId, request.SceneId, request.ServerId); ok {
		return nil, errors.New("session is already open")
	}

	log.Printf("opening session uid:%s sceneId:%s characterId:%s playerName:%s serverId%s\n",
		request.Uid, request.SceneId, request.CharacterId, request.PlayerName, request.ServerId)

	ses := api.NewSession(&goinworld.ClientConfig{
		Connection: &goinworld.ConnectionConfig{
			AutoReconnect:    false,
			DisconnectTimout: DisconnectTimeout,
		},
		Capabilities: &engine.CapabilitiesRequest{
			Audio:    false,
			Emotions: Emotions,
		},
	}, ApiKey, request.Uid, request.SceneId, request.CharacterId, request.PlayerName, request.ServerId)

	character := ses.GetCharacter()
	characters := ses.GetCharacters()
	chars := make([]*Character, len(characters))
	for i, s := range characters {
		chars[i] = CharacterFromSession(s)
	}

	return &OpenSessionResponse{
		SessionId:  ses.GetSessionId(),
		Character:  CharacterFromSession(character),
		Characters: chars,
	}, nil
}

func (a *Application) findSession(uid, characterId, sceneId, serverId string) (*api.Session, bool) {
	a.sessionMu.RLock()
	defer a.sessionMu.RUnlock()

	for _, ses := range a.Sessions {
		if ses.GetUid() == uid &&
			ses.GetCharacterId() == characterId &&
			ses.GetSceneId() == sceneId &&
			ses.GetServerId() == serverId { // also works for the case where it's not present
			return ses, true
		}
	}
	return nil, false
}
