package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api/v1")

	guestUsers := v1.Group("/users")
	guestUsers.Post("", h.SignUp)
}
