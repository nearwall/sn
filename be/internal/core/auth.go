package core

import (
	"context"
	"errors"

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

var (
	ErrLoginCreds = errors.New("no such account or wrong password")
)

type AuthService interface {
	LoginWithPassword(ctx context.Context, data PasswordLoginData) (PasswordLoginOk, error)
}
