package core

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	RegistrationData struct {
		Info     PersonalInfo
		Password string
	}

	CreationOk struct {
		UserID uuid.UUID
	}

	PersonalInfo struct {
		FirstName  string
		SecondName string
		Birthdate  time.Time
		Biography  string
		City       string
	}
)

var (
	ErrAccountExist    = errors.New("an account with this ID already exists")
	ErrAccountNotFound = errors.New("no account with this ID")
)

type AccountService interface {
	Create(ctx context.Context, data RegistrationData) (CreationOk, error)

	GetInfo(ctx context.Context, userID uuid.UUID) (PersonalInfo, error)
}

type InfoStore interface {
	LinkToAccount(ctx context.Context, accountID uuid.UUID, info PersonalInfo) error

	ReadInfo(ctx context.Context, accountID uuid.UUID) (PersonalInfo, error)
}

type (
	AccountCreationData struct {
		AccountID uuid.UUID
		Password  HashedPassword
	}

	PasswordInfo struct {
		Password  HashedPassword
		UpdatedAt time.Time
	}
)

type AccountStore interface {
	Create(ctx context.Context, data AccountCreationData) error

	ReadPasswordInfo(ctx context.Context, accountID uuid.UUID) (PasswordInfo, error)
}
