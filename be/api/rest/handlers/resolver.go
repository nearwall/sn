package handlers

import (
	"context"
	"io"
	"net/http"

	api "sn/api/rest/generated"
	"sn/internal/core"
	"sn/internal/infra/logger"
)

type (
	Resolver struct {
		user core.UserService
	}

	BearerTokenAuth struct {
	}

	Server struct {
		resolver  Resolver
		config    ServerConfig
		tokenAuth BearerTokenAuth
	}

	ServerConfig struct {
		Addr string
	}
)

func NewResolver(user core.UserService) Resolver {
	return Resolver{
		user: user,
	}
}

func NewServer(rslvr Resolver, config ServerConfig) Server {
	return Server{
		resolver:  rslvr,
		config:    config,
		tokenAuth: BearerTokenAuth{},
	}
}

func (srv Server) Run(ctx context.Context) error {
	server, err := api.NewServer(
		&srv.resolver,
		&srv.tokenAuth,
		api.WithNotFound(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			_, _ = io.WriteString(w, `{"error": "not found"}`)
		}))

	if err != nil {
		logger.Log().Fatalf(ctx, "fail to create server: {%s}", logger.ErrorLabel, err)
		return err
	}

	return http.ListenAndServe(srv.config.Addr, server)
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return nil
}

func (a *BearerTokenAuth) HandleBearerAuth(ctx context.Context, operationName api.OperationName, t api.BearerAuth) (context.Context, error) {
	return ctx, nil
}
