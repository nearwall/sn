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

	PersonalInfoEntity struct {
		PersonalInfo
		UserID uuid.UUID
	}

	SearchAccountsInfoParams struct {
		FirstName string
		LastName  string
		Limit     uint16
	}
)

var (
	ErrAccountExist    = errors.New("an account with this ID already exists")
	ErrAccountNotFound = errors.New("no account with this ID")
)

type AccountService interface {
	Create(ctx context.Context, data RegistrationData) (CreationOk, error)

	GetInfo(ctx context.Context, userID uuid.UUID) (PersonalInfoEntity, error)

	SearchAccounts(ctx context.Context, parameters SearchAccountsInfoParams) ([]PersonalInfoEntity, error)
}

type InfoStore interface {
	LinkToAccount(ctx context.Context, accountID uuid.UUID, info PersonalInfo) error

	ReadInfo(ctx context.Context, accountID uuid.UUID) (PersonalInfoEntity, error)

	GetInfoList(ctx context.Context, parameters SearchAccountsInfoParams) ([]PersonalInfoEntity, error)
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
