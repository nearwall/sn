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

// core.UserService interface
func (s *userService) Register(ctx context.Context, data core.RegistrationData) (core.RegistrationOk, error) {
	// generate random ID (UUID v4)
	userID := uuid.New()
	if err := s.storage.Create(ctx, userID); err != nil {
		return core.RegistrationOk{}, err
	}

	return core.RegistrationOk{
		UserID: userID,
	}, nil
}

// core.UserService interface
func (s *userService) GetInfo(ctx context.Context, userID uuid.UUID) (core.UserInfo, error) {
	return s.storage.ReadInfo(ctx, userID)
}
