package service

import (
	"context"
	"sn/internal/core"
)

type authService struct {
	storage  core.UserStore
	password core.PasswordService
}

func NewAuthService(storage core.UserStore, password core.PasswordService) core.AuthService {
	return &authService{
		storage: storage,
	}
}

// impl core.AuthService interface
func (s *authService) LoginWithPassword(ctx context.Context, data core.PasswordLoginData) (core.PasswordLoginOk, error) {
	return core.PasswordLoginOk{}, nil
}
