package schemes

import (
	"errors"

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

func FromLoginWithPasswordError(err error, reqID string) api.LoginPostRes {
	switch {
	case errors.Is(err, core.ErrLoginCreds):
		return &api.LoginPostBadRequest{}
	default:
		return &api.LoginPostInternalServerError{
			Response: api.R5xx{
				Message:   "Internal error",
				RequestID: api.NewOptString(reqID),
				Code:      api.OptInt{},
			},
			RetryAfter: api.OptInt{},
		}
	}
}

func FromLoginWithPasswordOk(data core.PasswordLoginOk) api.LoginPostRes {
	return &api.LoginPostOK{Token: api.NewOptString(data.AccessToken)}
}
