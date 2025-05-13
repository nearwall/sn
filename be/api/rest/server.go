package rest

import (
	"context"
	"io"
	"net/http"

	api "sn/api/rest/generated"
	"sn/api/rest/handlers"
	"sn/api/rest/middleware"

	"sn/internal/infra/logger"
)

type (
	Server struct {
		resolver  handlers.Resolver
		config    ServerConfig
		tokenAuth middleware.BearerTokenAuth
	}

	ServerConfig struct {
		Addr string
	}
)

func NewServer(rslvr handlers.Resolver, config ServerConfig) Server {
	return Server{
		resolver:  rslvr,
		config:    config,
		tokenAuth: middleware.BearerTokenAuth{},
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
