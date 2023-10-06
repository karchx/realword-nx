package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/karchx/realword-nx/model"
	"github.com/karchx/realword-nx/utils"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var u model.User
	req := &userRegisterRequest{}
	if err := req.bind(c, &u, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	if err := h.userStore.Create(&u); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newUserResponse(&u))
}

func (h *Handler) Login(c *fiber.Ctx) error {
	req := &userLoginRequest{}
	if err := req.bind(c, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	u, err := h.userStore.GetByEmail(req.User.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if u == nil {
		return c.Status(http.StatusForbidden).JSON(utils.AccessForbidden())
	}
	if !u.CheckPassword(req.User.Password) {
		return c.Status(http.StatusForbidden).JSON(utils.AccessForbidden())
	}

	return c.Status(http.StatusOK).JSON(newUserResponse(u))
}
