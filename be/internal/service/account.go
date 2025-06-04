package service

import (
	"context"

	"sn/internal/core"

	"github.com/google/uuid"
)

type accountService struct {
	infoStg  core.InfoStore
	accStg   core.AccountStore
	password core.PasswordService
}

func NewAccountService(info core.InfoStore, account core.AccountStore, password core.PasswordService) core.AccountService {
	return &accountService{
		infoStg:  info,
		accStg:   account,
		password: password,
	}
}

// core.UserService interface
func (s *accountService) Create(ctx context.Context, data core.RegistrationData) (core.CreationOk, error) {
	// random UUID v4
	accID := uuid.New()

	pwd, err := s.password.Hash(ctx, data.Password)
	if err != nil {
		return core.CreationOk{}, err
	}

	// FixMe: It should be a transaction with usrStg
	if err := s.accStg.Create(ctx, core.AccountCreationData{AccountID: accID, Password: pwd}); err != nil {
		return core.CreationOk{}, err
	}

	if err := s.infoStg.LinkToAccount(ctx, accID, data.Info); err != nil {
		return core.CreationOk{}, err
	}

	return core.CreationOk{
		UserID: accID,
	}, nil
}

// core.UserService interface
func (s *accountService) GetInfo(ctx context.Context, userID uuid.UUID) (core.PersonalInfo, error) {
	return s.infoStg.ReadInfo(ctx, userID)
}
