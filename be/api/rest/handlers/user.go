package handlers

import (
	"context"

	"github.com/google/uuid"

	api "sn/api/rest/generated"
	"sn/api/rest/handlers/schemes"
	"sn/internal/infra/logger"
)

// UserGetIDGet implements GET /user/get/{id} operation.
//
// Получение анкеты пользователя.
//
// GET /user/get/{id}
func (r *Resolver) UserGetIDGet(ctx context.Context, params api.UserGetIDGetParams) (api.UserGetIDGetRes, error) {
	logger.Log().Debug(ctx, "Handle POST /user/register")

	reqID, _ := ctx.Value(logger.RequestIDLabel).(string)

	userID, err := uuid.Parse(string(params.ID))
	if err != nil {
		logger.Log().Info(ctx, "incorrect user ID format", logger.ErrorLabel, err.Error())
		return &api.UserGetIDGetBadRequest{}, nil
	}

	info, err := r.user.GetInfo(ctx, userID)
	if err != nil {
		logger.Log().Info(ctx, "Fail to get user info", logger.ErrorLabel, err.Error())
		// It's a bit wired. User was deleted while we was handling this request
		return schemes.FromUserGetInfoErr(err, reqID), nil
	}

	return schemes.FromGetUserInfoOk(userID, info), nil
}

// UserRegisterPost implements POST /user/register operation.
//
// Регистрация нового пользователя.
//
// POST /user/register
func (r *Resolver) UserRegisterPost(ctx context.Context, req api.OptUserRegisterPostReq) (api.UserRegisterPostRes, error) {
	logger.Log().Debug(ctx, "Handle POST /user/register")

	reqID, _ := ctx.Value(logger.RequestIDLabel).(string)

	data, err := schemes.ToCoreRegistration(req)
	if err != nil {
		return &api.UserRegisterPostBadRequest{}, nil
	}

	success, err := r.user.Create(ctx, data)
	if err != nil {
		return schemes.FromRegistrationErr(err, reqID), nil
	}

	return schemes.FromRegistrationOk(success), nil
}

// UserSearchGet implements GET /user/search operation.
//
// Поиск анкет.
//
// GET /user/search
func (r *Resolver) UserSearchGet(ctx context.Context, params api.UserSearchGetParams) (api.UserSearchGetRes, error) {
	return &api.UserSearchGetServiceUnavailable{}, nil
}
