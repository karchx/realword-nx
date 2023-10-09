package handler

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/karchx/realword-nx/utils"
)

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api/v1")
	jwtMiddleware := jwtware.New(
		jwtware.Config{
			SigningKey: utils.JWTSecret,
			AuthScheme: "Bearer",
		})

	guestUsers := v1.Group("/users")
	guestUsers.Post("", h.SignUp)
	guestUsers.Post("/login", h.Login)
	user := v1.Group("/user", jwtMiddleware)
	user.Get("", h.CurrentUser)
}
