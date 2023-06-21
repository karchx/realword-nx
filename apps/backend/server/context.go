package server

import (
	"context"

	"github.com/karchx/realword-nx/conduit"
)

type contextKey string

const (
	userKey contextKey = "user"
)

func userFromContext(ctx context.Context) *conduit.User {
	user, ok := ctx.Value(userKey).(*conduit.User)

	if !ok {
		panic("missing user context key")
	}
	return user
}
