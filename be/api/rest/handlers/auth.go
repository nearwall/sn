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
	data, err := schemes.ToCoreLoginData(req)
	if err != nil {
		return &api.LoginPostBadRequest{}, nil
	}
	ctx = context.WithValue(ctx, logger.UserIDLabel, data.UserID)

	logger.Log().Debug(ctx, "Handle POST /login")

	if tokens, err := r.auth.LoginWithPassword(ctx, data); err != nil {
		return schemes.FromCoreLoginWithPasswordError(err), nil
	} else {
		return schemes.FromCoreLoginWithPasswordOk(tokens), nil
	}
}
