package schemes

import (
	"github.com/google/uuid"

	api "sn/api/rest/generated"
	"sn/internal/core"
)

func ToCoreLoginData(req api.OptLoginPostReq) (core.PasswordLoginData, error) {
	userID, err := uuid.Parse(string(req.Value.ID.Value))
	if err != nil {
		return core.PasswordLoginData{}, nil
	}

	return core.PasswordLoginData{
		UserID:   userID,
		Password: req.Value.Password.Value,
	}, nil
}

func FromCoreLoginWithPasswordError(err error) api.LoginPostRes {
	// LoginPostNotFound
	// LoginPostInternalServerError
	return &api.LoginPostInternalServerError{}

}

func FromCoreLoginWithPasswordOk(data core.PasswordLoginOk) api.LoginPostRes {
	return &api.LoginPostOK{Token: api.NewOptString(data.AccessToken)}
}
