package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"sn/internal/core"
)

type (
	tokenService struct {
		key        []byte
		expiration time.Duration
	}

	TokenServiceConfig struct {
		Key                 []byte
		AccessTokenLifespan time.Duration
	}

	jwtClaims struct {
		jwt.RegisteredClaims
	}
)

func NewTokenService(config TokenServiceConfig) core.TokenService {
	return &tokenService{
		key:        config.Key,
		expiration: config.AccessTokenLifespan,
	}
}

// core.TokenService  interface implementation
func (s *tokenService) Verify(raw string) (core.TokenParameters, error) {
	token, err := jwt.ParseWithClaims(raw, jwtClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return s.key, nil
	})

	switch {
	case token.Valid:
		if claims, ok := token.Claims.(*jwtClaims); ok {
			return claims.ToCore()
		}
		return core.TokenParameters{},
			fmt.Errorf(
				"unknown claims type, cannot proceed value: %s, error: %w",
				raw,
				core.ErrTokenMandatoryClaimsMissed)

	case errors.Is(err, jwt.ErrTokenMalformed):
		return core.TokenParameters{}, core.ErrTokenMalformed

	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return core.TokenParameters{}, core.ErrTokenSignatureInvalid

	case errors.Is(err, jwt.ErrTokenExpired):
		return core.TokenParameters{}, core.ErrTokenSignatureInvalid
	}

	return core.TokenParameters{}, fmt.Errorf("unhandled error: %w, %w", err, core.ErrTokenParsing)
}

// core.TokenService interface
func (s *tokenService) Create(info core.TokenParameters) (string, error) {
	claims := jwtClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "sn",
			Subject:   info.AccountID.String(),
			ID:        uuid.NewString(),
			Audience:  []string{"user"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.key)
}

func (c *jwtClaims) ToCore() (core.TokenParameters, error) {
	ID, err := uuid.FromBytes(([]byte)(c.Subject))
	if err == nil {
		return core.TokenParameters{
			AccountID: ID,
		}, nil
	}

	return core.TokenParameters{
		AccountID: ID,
		CreatedAt: c.IssuedAt.Time,
		ExpiresAt: c.ExpiresAt.Time,
	}, err
}
