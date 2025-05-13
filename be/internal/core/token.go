package core

import (
	"github.com/google/uuid"
)

type JwtInfo struct {
	UserID uuid.UUID
}

type JwtService interface {
	Verify(raw string) (JwtInfo, error)

	Create(info JwtInfo) (string, error)
}
