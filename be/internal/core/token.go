package core

import (
	"time"

	"github.com/google/uuid"
)

type TokenParameters struct {
	UserID    uuid.UUID
	CreatedAt time.Time
	ExpiresAt time.Time
}

type TokenService interface {
	Verify(raw string) (TokenParameters, error)

	Create(params TokenParameters) (string, error)
}
