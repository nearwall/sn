package schemes

import (
	api "sn/api/rest/generated"
	"sn/internal/core"

	"github.com/google/uuid"
)

func ToCoreRegistration(req api.OptUserRegisterPostReq) (core.RegistrationData, error) {
	return core.RegistrationData{}, nil
}

func FromRegistrationOk(data core.RegistrationOk) api.UserRegisterPostRes {
	return &api.UserRegisterPostOK{
		UserID: api.NewOptString(data.UserID.String()),
	}
}

func FromRegistrationErr(er error) api.UserRegisterPostRes {
	return &api.UserRegisterPostInternalServerError{}
}

func FromGetUserInfoOk(userID uuid.UUID, info core.UserInfo) api.UserGetIDGetRes {
	return &api.User{
		ID:         api.NewOptUserId(api.UserId(userID.String())),
		FirstName:  api.NewOptString(info.FirstName),
		SecondName: api.NewOptString(info.SecondName),
		Biography:  api.NewOptString(info.Biography),
		Birthdate:  api.NewOptBirthDate(api.BirthDate(info.Birthdate)),
		City:       api.NewOptString(info.City),
	}
}
