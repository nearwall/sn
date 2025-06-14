package service

import (
	"context"
	"sort"

	"github.com/google/uuid"

	"sn/internal/core"
)

type accountService struct {
	infoStg  core.InfoStore
	accStg   core.AccountStore
	password core.PasswordService
}

func NewAccountService(
	info core.InfoStore,
	account core.AccountStore,
	password core.PasswordService,
) core.AccountService {
	return &accountService{
		infoStg:  info,
		accStg:   account,
		password: password,
	}
}

// core.AccountService interface
func (s *accountService) Create(ctx context.Context, data core.RegistrationData) (core.CreationOk, error) {
	// random UUID v4
	accID := uuid.New()

	pwd, err := s.password.Hash(ctx, data.Password)
	if err != nil {
		return core.CreationOk{}, err
	}

	// FixMe: It should be a transaction with usrStg
	if err = s.accStg.Create(ctx, core.AccountCreationData{AccountID: accID, Password: pwd}); err != nil {
		return core.CreationOk{}, err
	}

	if err = s.infoStg.LinkToAccount(ctx, accID, data.Info); err != nil {
		return core.CreationOk{}, err
	}

	return core.CreationOk{
		UserID: accID,
	}, nil
}

// core.AccountService interface
func (s *accountService) GetInfo(ctx context.Context, userID uuid.UUID) (core.PersonalInfoEntity, error) {
	return s.infoStg.ReadInfo(ctx, userID)
}

// core.AccountService interface
func (s *accountService) SearchAccounts(ctx context.Context, parameters core.SearchAccountsInfoParams) ([]core.PersonalInfoEntity, error) {
	people, err := s.infoStg.GetInfoList(ctx, parameters)
	if err != nil {
		return []core.PersonalInfoEntity{}, err
	}

	sort.Slice(people, func(i, j int) bool {
		for index := range len(people[i].UserID) {
			if people[i].UserID[index] < people[j].UserID[index] {
				return true
			} else if people[i].UserID[index] < people[j].UserID[index] {
				return false
			}
		}
		return false
	})

	return people, nil
}
