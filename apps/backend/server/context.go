package server

import (
	"context"
	"net/http"

	"github.com/karchx/realword-nx/conduit"
)

type contextKey string

const (
	userKey  contextKey = "user"
	tokenKey contextKey = "token"
)

func setContextUser(r *http.Request, u *conduit.User) *http.Request {
	ctx := context.WithValue(r.Context(), userKey, u)
	return r.WithContext(ctx)
}

func userFromContext(ctx context.Context) *conduit.User {
	user, ok := ctx.Value(userKey).(*conduit.User)

	if !ok {
		panic("missing user context key")
	}
	return user
}

func setContextUserToken(r *http.Request, token string) *http.Request {
	ctx := context.WithValue(r.Context(), tokenKey, token)
	return r.WithContext(ctx)
}
