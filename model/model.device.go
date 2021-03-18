package model

import (
	uuid "github.com/satori/go.uuid"
)

const (
	ROLE_MEDIA_CONTROLLER = 0
	ROLE_MEDIA_PLAYER     = 1
)

type (
	Device struct {
		ID     string    `json:"id"`
		UserID uuid.UUID `json:"user_id"`
		Name   string    `json:"name"`
		Role   int       `json:"role"`
	}
)
