package service

import (
	"errors"
	"fmt"
	"log"
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
		Key []byte
	}

	jwtClaims struct {
		jwt.RegisteredClaims
	}
)

// FixMe: use config
const EXPIRATION_HOURS = 24

func NewTokenService(config TokenServiceConfig) core.TokenService {
	return &tokenService{
		key:        config.Key,
		expiration: EXPIRATION_HOURS * time.Hour}
}

// core.TokenService  interface implementation
func (s *tokenService) Verify(raw string) (core.TokenParameters, error) {
	token, err := jwt.ParseWithClaims(raw, jwtClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return s.key, nil
	})

	switch {
	case token.Valid:
		fmt.Println("valid")
		if claims, ok := token.Claims.(*jwtClaims); ok {
			return claims.ToCore()
		}
		log.Fatal("unknown claims type, cannot proceed")
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

	return core.TokenParameters{}, nil
}

// core.TokenService interface
func (s *tokenService) Create(info core.TokenParameters) (string, error) {
	claims := jwtClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "sn",
			Subject:   info.UserID.String(),
			ID:        uuid.NewString(),
			Audience:  []string{"user"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.key)
}

func (c *jwtClaims) ToCore() (core.TokenParameters, error) {
	if userID, err := uuid.FromBytes(([]byte)(c.Subject)); err == nil {
		return core.TokenParameters{
			UserID: userID,
		}, nil
	} else {
		return core.TokenParameters{}, err
	}
}
