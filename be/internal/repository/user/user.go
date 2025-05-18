package repository

import (
	"context"

	"sn/internal/core"
	"sn/internal/infra/postgres"

	"github.com/google/uuid"
)

type userStore struct {
	client *postgres.Client
}

func NewUserStore(postgresClient *postgres.Client) core.UserStore {
	return &userStore{
		client: postgresClient,
	}
}

// core.UserStore interface
func (s *userStore) Create(ctx context.Context, userID uuid.UUID) error {
	return nil
}

// core.UserStore interface
func (s *userStore) ReadInfo(ctx context.Context, userID uuid.UUID) (core.UserInfo, error) {
	return core.UserInfo{}, nil
}
