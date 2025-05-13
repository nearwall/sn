package handlers

import (
	"sn/internal/core"
)

type (
	Resolver struct {
		user core.UserService
	}
)

func NewResolver(user core.UserService) Resolver {
	return Resolver{
		user: user,
	}
}
