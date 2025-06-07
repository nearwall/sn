package rest

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	api "sn/api/rest/generated"
	"sn/api/rest/handlers"
	"sn/api/rest/middleware"
	"sn/internal/infra/logger"
)

type (
	Server struct {
		resolver        handlers.Resolver
		config          ServerConfig
		tokenAuth       middleware.BearerTokenAuth
		shutdownChannel chan struct{}
		// Todo add `metrics` for api.WithMeterProvider
		// ToDo: add `tracing` for api.WithTracerProvider
	}

	ServerConfig struct {
		Addr string
	}
)

func NewServer(rslvr handlers.Resolver, tokenAuth middleware.BearerTokenAuth, config ServerConfig) Server {
	return Server{
		resolver:        rslvr,
		config:          config,
		tokenAuth:       tokenAuth,
		shutdownChannel: make(chan struct{}),
	}
}

func (srv *Server) Run(ctx context.Context) error {
	server, err := api.NewServer(
		&srv.resolver,
		&srv.tokenAuth,
		api.WithNotFound(func(w http.ResponseWriter, r *http.Request) {
			logger.Log().Info(ctx, "Unhandled request: %s %s", r.Method, r.URL.Path)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			_, _ = io.WriteString(w, `{"error": "not found"}`)
		}),
	)

	if err != nil {
		logger.Log().Fatalf(ctx, "fail to create server: {%s}", logger.ErrorLabel, err)
		return err
	}

	httpServer := http.Server{
		ReadHeaderTimeout: time.Second,
		Addr:              srv.config.Addr,
		Handler: middleware.Wrap(server,
			middleware.LogRequests(),
		),
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var shutdownReason string
		select {
		// Wait for shutdown request
		case <-srv.shutdownChannel:
			shutdownReason = "shutdown demand"
		// Wait until g ctx canceled
		case <-ctx.Done():
			shutdownReason = "root context cancellation"
		}

		shutdownCtx := context.Background()
		logger.Log().Infof(shutdownCtx, "Shutting down because of %s", shutdownReason)

		return httpServer.Shutdown(shutdownCtx)
	})

	g.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	return g.Wait()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	close(srv.shutdownChannel)
	return nil
}
