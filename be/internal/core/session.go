package core

import (
	"context"

	"github.com/google/uuid"
)

type OpenedSessionData struct {
	RawAccessToken string
}

type SessionService interface {
	Open(ctx context.Context, UserID uuid.UUID) (OpenedSessionData, error)

	Close(ctx context.Context, UserID uuid.UUID) (bool, error)
}
