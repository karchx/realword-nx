package server

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
)

func Logger(w io.Writer) func(h http.Handler) http.Handler {
	return (func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(w, h)
	})
}

func (s *Server) authenticate(mustAuth bool) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Vary", "Authorization")
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				if mustAuth {
					invalidAuthTokenError(w)
				}
			}
		})
	}
}
