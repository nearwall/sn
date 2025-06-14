package schemes

import (
	"errors"
	api "sn/api/rest/generated"
	"sn/internal/core"
	"time"
)

const UserSearchEntitiesLimit uint16 = 1000

func ToCoreRegistration(req api.OptUserRegisterPostReq) (core.RegistrationData, error) {
	birthDate := time.Time(req.Value.Birthdate.Value)

	return core.RegistrationData{
		Info: core.PersonalInfo{
			FirstName:  &req.Value.FirstName.Value,
			SecondName: &req.Value.SecondName.Value,
			Birthdate:  &birthDate,
			City:       &req.Value.City.Value,
			Biography:  &req.Value.Biography.Value},
		Password: req.Value.Password.Value,
	}, nil
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

func FromGetUserInfoOk(info core.PersonalInfoEntity) api.UserGetIDGetRes {
	user := fromCorePersonalInfoEntityToUser(info)
	return &user
}

func fromCorePersonalInfoEntityToUser(info core.PersonalInfoEntity) api.User {
	resp := api.User{
		ID: api.NewOptUserId(api.UserId(info.UserID.String())),
	}

	if info.FirstName != nil {
		resp.FirstName = api.NewOptString(*info.FirstName)
	}
	if info.SecondName != nil {
		resp.SecondName = api.NewOptString(*info.SecondName)
	}
	if info.Biography != nil {
		resp.Biography = api.NewOptString(*info.Biography)
	}
	if info.Birthdate != nil {
		resp.Birthdate = api.NewOptBirthDate(api.BirthDate(*info.Birthdate))
	}
	if info.City != nil {
		resp.City = api.NewOptString(*info.City)
	}

	return resp
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

func FromUserSearchOk(data []core.PersonalInfoEntity) api.UserSearchGetRes {
	var users api.UserSearchGetOKApplicationJSON = make([]api.User, 0, len(data))

	for _, value := range data {
		users = append(users, fromCorePersonalInfoEntityToUser(value))
	}

	return &users
}

func FromUserSearchErr(err error, reqID string) api.UserSearchGetRes {
	switch {
	case errors.Is(err, core.ErrAccountExist):
		return &api.UserSearchGetBadRequest{}
	default:
		return &api.UserSearchGetInternalServerError{
			Response: api.R5xx{
				Message:   "Internal error",
				RequestID: api.NewOptString(reqID),
				Code:      api.OptInt{},
			},
			RetryAfter: api.OptInt{},
		}
	}
}
