package goinworld

import (
	"errors"
	"github.com/zivoy/go-inworld/internal/protoBuf/engine"
)

type InworldClient struct {
	apiKey *ApiKey
	user   *engine.UserRequest
	scene  string
	client *engine.ClientRequest
	config *ClientConfiguration

	generateSessionTokenFunc GenerateSessionTokenFunc
}

func (a *ApiKey) Validate() error {
	if a == nil || a.Key == "" || a.Secret == "" {
		return errors.New("api key is required")
	}
	return nil
}

func (c *InworldClient) Validate() error {
	if c.generateSessionTokenFunc == nil {
		return c.apiKey.Validate()
	}

	if c.scene != "" {
		return errors.New("scene name is required")
	}
	return nil
}
