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
	user.Put("", h.UpdateUser)

	profiles := v1.Group("/profiles", jwtMiddleware)
	profiles.Get("/:username", h.GetProfile)
	profiles.Post("/:username/follow", h.Follow)
	profiles.Delete("/:username/unfollow", h.UnFollow)

	articles := v1.Group("/articles", jwtMiddleware)
	articles.Post("", h.CreateArticle)
	articles.Put("/:slug", h.UpdateArticle)
	articles.Get("/:slug", h.GetArticle)
}
