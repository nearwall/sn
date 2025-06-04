package handlers

import (
	"sn/internal/core"
)

type (
	Resolver struct {
		user core.AccountService
		auth core.AuthService
	}
)

func NewResolver(user core.AccountService, auth core.AuthService) Resolver {
	return Resolver{
		user: user,
		auth: auth,
	}
}
