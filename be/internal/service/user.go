package service

import (
	"context"
	"sn/internal/core"

	"github.com/google/uuid"
)

type service struct {
	storage core.UserStore
}

func NewService(storage core.UserStore) core.UserService {
	return &service{
		storage: storage,
	}
}

func (s *service) Register(ctx context.Context) (uuid.UUID, error) {
	// generate random ID (UUID v4)
	user_id := uuid.New()
	if err := s.storage.Create(ctx, user_id); err != nil {
		return uuid.UUID{}, err
	}

	return user_id, nil
}

func (s *service) GetInfo(ctx context.Context, user_id uuid.UUID) (*core.UserInfo, error) {
	return s.storage.ReadInfo(ctx, user_id)
}
