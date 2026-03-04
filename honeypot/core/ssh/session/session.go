package session

import (
	"fmt"
	"honeypot/core/config"

	"github.com/google/uuid"
)

type Config struct {
	Config *config.Config
}

func NewID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("UUID Generate Error: %s", err.Error())
	}

	return id.String(), nil
}
