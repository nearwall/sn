package handlers

import (
	"context"

	api "sn/api/rest/generated"
	"sn/api/rest/handlers/schemes"
	"sn/internal/infra/logger"
)

// LoginPost implements POST /login operation.
//
// Упрощенный процесс аутентификации путем передачи
// идентификатор пользователя и получения токена для
// дальнейшего прохождения авторизации.
//
// POST /login
//
// api.Handler interface implementation
func (r *Resolver) LoginPost(ctx context.Context, req api.OptLoginPostReq) (api.LoginPostRes, error) {
	reqID, _ := ctx.Value(logger.RequestIDLabel).(string)

	data, err := schemes.ToCoreLoginData(req)
	if err != nil {
		//nolint: nilerr // api.LoginPostBadRequest - is result for such errors
		return &api.LoginPostBadRequest{}, nil
	}
	ctx = context.WithValue(ctx, logger.UserIDLabel, data.UserID)

	logger.Log().Debug(ctx, "Handle POST /login")

	tokens, err := r.auth.LoginWithPassword(ctx, data)
	if err != nil {
		return schemes.FromLoginWithPasswordError(err, reqID), nil
	}

	return schemes.FromLoginWithPasswordOk(tokens), nil
}
