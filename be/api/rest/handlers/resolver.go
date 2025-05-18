package handlers

import (
	"sn/internal/core"
)

type (
	Resolver struct {
		user core.UserService
		auth core.AuthService
	}
)

func NewResolver(user core.UserService, auth core.AuthService) Resolver {
	return Resolver{
		user: user,
		auth: auth,
	}
}
