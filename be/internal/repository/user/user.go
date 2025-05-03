package repository

import (
	"context"

	"sn/internal/core"
	"sn/internal/infra/postgres"

	"github.com/google/uuid"
)

type store struct {
	client *postgres.Client
}

func NewStore(postgresClient *postgres.Client) core.UserStore {
	return &store{
		client: postgresClient,
	}
}

// core.UserStore interface
func (s *store) Create(ctx context.Context, user_id uuid.UUID) error {
	return nil
}

// core.UserStore interface
func (s *store) ReadInfo(ctx context.Context, user_id uuid.UUID) (*core.UserInfo, error) {
	return &core.UserInfo{}, nil
}
