package schemes

import (
	api "sn/api/rest/generated"
	"sn/internal/core"

	"github.com/google/uuid"
)

func ConvertToUserInfo(userID uuid.UUID, info *core.UserInfo) api.UserGetIDGetRes {
	return &api.User{
		ID:         api.NewOptUserId(api.UserId(userID.String())),
		FirstName:  api.NewOptString(info.FirstName),
		SecondName: api.NewOptString(info.SecondName),
		Biography:  api.NewOptString(info.Biography),
		Birthdate:  api.NewOptBirthDate(api.BirthDate(info.Birthdate)),
		City:       api.NewOptString(info.City),
	}
}

func ConvertToCreatedUserID(userID uuid.UUID) api.UserRegisterPostRes {
	return &api.UserRegisterPostOK{
		UserID: api.NewOptString(userID.String()),
	}
}
