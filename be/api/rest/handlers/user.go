package handlers

import (
	"context"

	api "sn/api/rest/generated"
	"sn/api/rest/handlers/schemes"
	"sn/internal/infra/logger"

	"github.com/google/uuid"
)

// UserGetIDGet implements GET /user/get/{id} operation.
//
// Получение анкеты пользователя.
//
// GET /user/get/{id}
func (r *Resolver) UserGetIDGet(ctx context.Context, params api.UserGetIDGetParams) (api.UserGetIDGetRes, error) {
	userID, err := uuid.Parse(string(params.ID))
	if err != nil {
		logger.Log().Info(ctx, "incorrect user ID format", logger.ErrorLabel, err.Error())
		return &api.UserGetIDGetBadRequest{}, nil
	}

	reqID, _ := ctx.Value(logger.RequestIDLabel).(string)

	info, err := r.user.GetInfo(ctx, userID)
	if err != nil {
		logger.Log().Info(ctx, "Fail to get user info", logger.ErrorLabel, err.Error())
		return &api.UserGetIDGetInternalServerError{
			Response: api.R5xx{
				Message:   "Internal error",
				RequestID: api.NewOptString(reqID),
				Code:      api.OptInt{},
			},
			RetryAfter: api.OptInt{},
		}, nil
	}
	if info == nil {
		// It's a bit wired. User was deleted while we was handling this request
		logger.Log().Error(ctx, "No user info found")
		return &api.UserGetIDGetNotFound{}, nil
	}

	return schemes.FromGetUserInfoOk(userID, info), nil
}

// UserRegisterPost implements POST /user/register operation.
//
// Регистрация нового пользователя.
//
// POST /user/register
func (r *Resolver) UserRegisterPost(ctx context.Context, req api.OptUserRegisterPostReq) (api.UserRegisterPostRes, error) {
	userID, err := r.user.Register(ctx)
	if err != nil {
		return nil, err
	}

	return schemes.ConvertToCreatedUserID(userID), nil
}

// UserSearchGet implements GET /user/search operation.
//
// Поиск анкет.
//
// GET /user/search
func (r *Resolver) UserSearchGet(ctx context.Context, params api.UserSearchGetParams) (api.UserSearchGetRes, error) {
	return nil, nil
}
