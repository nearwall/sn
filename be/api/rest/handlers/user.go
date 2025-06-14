package handlers

import (
	"context"

	"github.com/google/uuid"

	api "sn/api/rest/generated"
	"sn/api/rest/handlers/schemes"
	"sn/internal/core"
	"sn/internal/infra/logger"
)

// UserGetIDGet implements GET /user/get/{id} operation.
//
// Получение анкеты пользователя.
//
// GET /user/get/{id}
func (r *Resolver) UserGetIDGet(ctx context.Context, params api.UserGetIDGetParams) (api.UserGetIDGetRes, error) {
	reqID, _ := ctx.Value(logger.RequestIDLabel).(string)

	userID, err := uuid.Parse(string(params.ID))
	if err != nil {
		logger.Log().Info(ctx, "incorrect user ID format", logger.ErrorLabel, err.Error())
		//nolint: nilerr // response UserGetIDGetBadRequest will be returned instead of error
		return &api.UserGetIDGetBadRequest{}, nil
	}

	info, err := r.user.GetInfo(ctx, userID)
	if err != nil {
		logger.Log().Info(ctx, "Fail to get user info", logger.ErrorLabel, err.Error())
		// It's a bit wired. User was deleted while we was handling this request
		//nolint: nilerr // corresponding error response will be returned instead of error
		return schemes.FromUserGetInfoErr(err, reqID), nil
	}

	logger.Log().Infof(ctx, "account %s info was read successfully", userID)
	return schemes.FromGetUserInfoOk(userID, info), nil
}

// UserRegisterPost implements POST /user/register operation.
//
// Регистрация нового пользователя.
//
// POST /user/register
func (r *Resolver) UserRegisterPost(ctx context.Context, req api.OptUserRegisterPostReq) (api.UserRegisterPostRes, error) {
	reqID, _ := ctx.Value(logger.RequestIDLabel).(string)

	data, err := schemes.ToCoreRegistration(req)
	if err != nil {
		logger.Log().Infof(ctx, "fail to parse request body: %w", err)
		//nolint: nilerr // response `UserRegisterPostBadRequest` will be returned instead of error
		return &api.UserRegisterPostBadRequest{}, nil
	}

	success, err := r.user.Create(ctx, data)
	if err != nil {
		logger.Log().Infof(ctx, "fail to create account: %w", err)
		//nolint: nilerr // corresponding error response will be returned instead of error
		return schemes.FromRegistrationErr(err, reqID), nil
	}

	logger.Log().Infof(ctx, "account %s was created successfully", success.UserID)
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
