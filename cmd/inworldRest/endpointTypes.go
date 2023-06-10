package main

import "github.com/zivoy/go-inworld/session"

type OpenSessionRequest struct {
	Uid         string `json:"uid" binding:"required"`
	SceneId     string `json:"sceneId" binding:"required"`
	CharacterId string `json:"characterId" binding:"required"`
	PlayerName  string `json:"playerName,omitempty"`
	ServerId    string `json:"serverId,omitempty"`
}

type OpenSessionResponse struct {
	SessionId  string       `json:"sessionId"`
	Character  *Character   `json:"character"`
	Characters []*Character `json:"characters"`
}

type Character struct {
	Id           string `json:"id"`
	ResourceName string `json:"resourceName"`
	DisplayName  string `json:"displayName"`
}

func CharacterFromSession(char *session.Character) *Character {
	return &Character{
		Id:           char.Id,
		ResourceName: char.ResourceName,
		DisplayName:  char.DisplayName,
	}
}
