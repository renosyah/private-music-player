package model

import uuid "github.com/satori/go.uuid"

type (
	UserAuth struct {
		SessionID uuid.UUID `json:"session_id"`
		User      `json:"user"`
	}
)
