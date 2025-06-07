package middleware

import (
	"context"
	"net/http"
	"sn/internal/infra/logger"
)

func LogRequests() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), logger.RequestIDLabel, "")

			logger.Log().Infof(ctx, "Request handling: %s", LabelHTTPMethod, r.Method, LabelHTTPPath, r.URL.Path)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
