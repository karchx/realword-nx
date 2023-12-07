package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/karchx/realword-nx/model"
	"github.com/karchx/realword-nx/utils"
	uuid "github.com/satori/go.uuid"
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

func (h *Handler) CurrentUser(c *fiber.Ctx) error {
	u, err := h.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if u == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}
	return c.Status(http.StatusOK).JSON(newUserResponse(u))
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	u, err := h.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if u == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}
	req := newUserUpdateRequest()
	req.populate(u)
	if err := req.bind(c, u, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	if err := h.userStore.Update(u); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newUserResponse(u))
}

func (h *Handler) GetProfile(c *fiber.Ctx) error {
	username := c.Params("username")
	u, err := h.userStore.GetByUsername(username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if u == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}
	return c.Status(http.StatusOK).JSON(newProfileResponse(h.userStore, userIDFromToken(c), u))
}

func (h *Handler) Follow(c *fiber.Ctx) error {
	followerID := userIDFromToken(c)
	username := c.Params("username")
	u, err := h.userStore.GetByUsername(username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if u == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}
	if err := h.userStore.AddFollower(u, followerID); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newProfileResponse(h.userStore, userIDFromToken(c), u))
}

func (h *Handler) UnFollow(c *fiber.Ctx) error {
	followerID := userIDFromToken(c)
	username := c.Params("username")
	u, err := h.userStore.GetByUsername(username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if u == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}
	if err := h.userStore.RemoveFollower(u, followerID); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newProfileResponse(h.userStore, userIDFromToken(c), u))
}

func userIDFromToken(c *fiber.Ctx) uuid.UUID {
	var user *jwt.Token
	l := c.Locals("user")
	if l == nil {
		return uuid.UUID{}
	}

	user = l.(*jwt.Token)
	id := fmt.Sprint(user.Claims.(jwt.MapClaims)["id"])
	uuidUser, _ := uuid.FromString(id)
	return uuidUser
}
