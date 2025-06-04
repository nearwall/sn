package service

import (
	"context"
	"fmt"
	"time"

	"sn/internal/core"

	"github.com/google/uuid"
)

type (
	sessionService struct {
		sessionDuration time.Duration
		token           core.TokenService
	}

	SessionServiceConfig struct {
		SessionDuration time.Duration
	}
)

func NewSessionService(config SessionServiceConfig, token core.TokenService) core.SessionService {
	return &sessionService{
		sessionDuration: config.SessionDuration,
		token:           token,
	}
}

// core SessionService interface
func (s *sessionService) Open(ctx context.Context, UserID uuid.UUID) (core.OpenedSessionData, error) {
	now := time.Now()

	if rawAccessToken, err := s.token.Create(core.TokenParameters{UserID: UserID, CreatedAt: now, ExpiresAt: now.Add(s.sessionDuration)}); err != nil {
		return core.OpenedSessionData{}, nil
	} else {
		return core.OpenedSessionData{RawAccessToken: rawAccessToken}, nil
	}
}

// core SessionService interface
func (s *sessionService) Close(ctx context.Context, UserID uuid.UUID) (bool, error) {
	return false, fmt.Errorf("not implemented")
}
