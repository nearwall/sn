package core

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	Register(ctx context.Context, data RegistrationData) (RegistrationOk, error)

	GetInfo(ctx context.Context, userID uuid.UUID) (UserInfo, error)
}

type (
	RegistrationData struct {
		UserInfo
		Password string
	}

	RegistrationOk struct {
		UserID uuid.UUID
	}

	UserInfo struct {
		FirstName  string
		SecondName string
		Birthdate  time.Time
		Biography  string
		City       string
	}
)

type UserStore interface {
	Create(ctx context.Context, userID uuid.UUID) error

	ReadInfo(ctx context.Context, userID uuid.UUID) (UserInfo, error)
}
