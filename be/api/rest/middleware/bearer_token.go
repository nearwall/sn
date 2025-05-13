package middleware

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"

	api "sn/api/rest/generated"
	"sn/internal/core"
	"sn/internal/infra/logger"
)

type (
	BearerTokenAuth struct {
		jwtHandler core.JwtService
	}
)

func NewBearerTokenAuth(jwtHandler core.JwtService) BearerTokenAuth {
	return BearerTokenAuth{jwtHandler: jwtHandler}
}

// api.SecurityHandler interface implementation
func (a *BearerTokenAuth) HandleBearerAuth(ctx context.Context, operationName api.OperationName, t api.BearerAuth) (context.Context, error) {
	if claims, err := a.jwtHandler.Verify(t.GetToken()); err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			fmt.Println("That's not even a token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			// Invalid signature
			fmt.Println("Invalid signature")
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		default:
			fmt.Println("Couldn't handle this token:", err)
		}
		return ctx, err
	} else {
		return context.WithValue(ctx, logger.UserIDLabel, claims.UserID), nil
	}
}
