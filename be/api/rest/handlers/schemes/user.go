package schemes

import (
	"errors"
	api "sn/api/rest/generated"
	"sn/internal/core"

	"github.com/google/uuid"
)

func ToCoreRegistration(req api.OptUserRegisterPostReq) (core.RegistrationData, error) {
	return core.RegistrationData{}, nil
}

func FromRegistrationOk(data core.CreationOk) api.UserRegisterPostRes {
	return &api.UserRegisterPostOK{
		UserID: api.NewOptString(data.UserID.String()),
	}
}

func FromRegistrationErr(err error, reqID string) api.UserRegisterPostRes {
	switch {
	case errors.Is(err, core.ErrAccountExist):
		return &api.UserRegisterPostBadRequest{}
	default:
		return &api.UserRegisterPostInternalServerError{
			Response: api.R5xx{
				Message:   "Internal error",
				RequestID: api.NewOptString(reqID),
				Code:      api.OptInt{},
			},
			RetryAfter: api.OptInt{},
		}
	}
}

func FromGetUserInfoOk(userID uuid.UUID, info core.PersonalInfo) api.UserGetIDGetRes {
	return &api.User{
		ID:         api.NewOptUserId(api.UserId(userID.String())),
		FirstName:  api.NewOptString(info.FirstName),
		SecondName: api.NewOptString(info.SecondName),
		Biography:  api.NewOptString(info.Biography),
		Birthdate:  api.NewOptBirthDate(api.BirthDate(info.Birthdate)),
		City:       api.NewOptString(info.City),
	}
}

func FromUserGetInfoErr(err error, reqID string) api.UserGetIDGetRes {
	switch {
	case errors.Is(err, core.ErrAccountNotFound):
		return &api.UserGetIDGetNotFound{}
	default:
		return &api.UserGetIDGetInternalServerError{
			Response: api.R5xx{
				Message:   "Internal error",
				RequestID: api.NewOptString(reqID),
				Code:      api.OptInt{},
			},
			RetryAfter: api.OptInt{},
		}
	}
}
