package handlers

import (
	"context"

	api "sn/api/rest/generated"
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
	//
	logger.Log().Debug(ctx, "Handle POST /login")

	return &api.LoginPostBadRequest{}, nil
}
