package service

import (
	"context"
	"fmt"

	"sn/internal/core"
)

type authService struct {
	password core.PasswordService
	session  core.SessionService
	userStg  core.InfoStore
	accStg   core.AccountStore
}

func NewAuthService(
	storage core.InfoStore,
	accStg core.AccountStore,
	password core.PasswordService,
	session core.SessionService) core.AuthService {
	return &authService{
		userStg:  storage,
		accStg:   accStg,
		password: password,
		session:  session,
	}
}

// impl core.AuthService interface
func (a *authService) LoginWithPassword(ctx context.Context, data core.PasswordLoginData) (core.PasswordLoginOk, error) {
	pwdInfo, err := a.accStg.ReadPasswordInfo(ctx, data.UserID)
	if err != nil {
		return core.PasswordLoginOk{}, err
	}

	if isCorrect, err := a.password.Verify(ctx, data.Password, pwdInfo.Password); err != nil {
		return core.PasswordLoginOk{}, err
	} else if !isCorrect {
		return core.PasswordLoginOk{}, fmt.Errorf("Wrong password: %w", core.ErrLoginCreds)
	}

	sessionData, err := a.session.Open(ctx, data.UserID)
	if err != nil {
		return core.PasswordLoginOk{}, err
	}

	return core.PasswordLoginOk{AccessToken: sessionData.RawAccessToken}, nil
}
