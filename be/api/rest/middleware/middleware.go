package middleware

import "net/http"

type LogTagsHTTP = string

const (
	LabelHTTPMethod  LogTagsHTTP = "http_method"
	LabelHTTPPath    LogTagsHTTP = "http_path"
	LabelHTTPHandler LogTagsHTTP = "http_handler"
)

type Middleware = func(http.Handler) http.Handler

// Wrap handler using given middlewares.
func Wrap(h http.Handler, middlewares ...Middleware) http.Handler {
	switch len(middlewares) {
	case 0:
		return h
	case 1:
		return middlewares[0](h)
	default:
		for i := len(middlewares) - 1; i >= 0; i-- {
			h = middlewares[i](h)
		}
		return h
	}
}
