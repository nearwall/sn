package core

import (
	"context"

	"github.com/google/uuid"
)

type (
	PasswordLoginData struct {
		UserID   uuid.UUID
		Password string
	}

	PasswordLoginOk struct {
		AccessToken string
	}
)

type AuthService interface {
	LoginWithPassword(ctx context.Context, data PasswordLoginData) (PasswordLoginOk, error)
}
