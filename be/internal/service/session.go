package service

import (
	"context"
	"errors"
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
func (s *sessionService) Open(_ context.Context, userID uuid.UUID) (core.OpenedSessionData, error) {
	now := time.Now()
	rawAccessToken, err := s.token.Create(core.TokenParameters{AccountID: userID, CreatedAt: now, ExpiresAt: now.Add(s.sessionDuration)})
	if err != nil {
		return core.OpenedSessionData{}, err
	}

	return core.OpenedSessionData{RawAccessToken: rawAccessToken}, nil
}

// core SessionService interface
func (s *sessionService) Close(_ context.Context, _UserID uuid.UUID) (bool, error) {
	return false, errors.New("not implemented")
}
