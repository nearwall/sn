package core

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TokenParameters struct {
	AccountID uuid.UUID
	CreatedAt time.Time
	ExpiresAt time.Time
}

var (
	ErrTokenMandatoryClaimsMissed = errors.New("no mandatory claims")
	ErrTokenMalformed             = errors.New("token format isn't correct")
	ErrTokenSignatureInvalid      = errors.New("token signature is wrong")
	ErrTokenExpired               = errors.New("token expired")
	ErrTokenParsing               = errors.New("token parsing common error")
)

type TokenService interface {
	Verify(raw string) (TokenParameters, error)

	Create(params TokenParameters) (string, error)
}
