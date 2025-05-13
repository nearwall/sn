package service

import (
	"context"
	"sn/internal/core"

	"github.com/google/uuid"
)

type userService struct {
	storage core.UserStore
}

func NewUserService(storage core.UserStore) core.UserService {
	return &userService{
		storage: storage,
	}
}

func (s *userService) Register(ctx context.Context) (uuid.UUID, error) {
	// generate random ID (UUID v4)
	userID := uuid.New()
	if err := s.storage.Create(ctx, userID); err != nil {
		return uuid.UUID{}, err
	}

	return userID, nil
}

func (s *userService) GetInfo(ctx context.Context, userID uuid.UUID) (*core.UserInfo, error) {
	return s.storage.ReadInfo(ctx, userID)
}
