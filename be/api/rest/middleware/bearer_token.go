package middleware

import (
	"context"

	api "sn/api/rest/generated"
	"sn/internal/core"
	"sn/internal/infra/logger"
)

type (
	BearerTokenAuth struct {
		jwtHandler core.TokenService
	}
)

func NewBearerTokenAuth(jwtHandler core.TokenService) BearerTokenAuth {
	return BearerTokenAuth{jwtHandler: jwtHandler}
}

// api.SecurityHandler interface implementation
func (a *BearerTokenAuth) HandleBearerAuth(ctx context.Context, operationName api.OperationName, t api.BearerAuth) (context.Context, error) {
	claims, err := a.jwtHandler.Verify(t.GetToken())
	if err != nil {
		logger.Log().Infof(ctx, "Validation failed: %s", err.Error(), LabelHTTPHandler, operationName)
		return ctx, err
	}

	return context.WithValue(ctx, logger.UserIDLabel, claims.AccountID), nil
}
