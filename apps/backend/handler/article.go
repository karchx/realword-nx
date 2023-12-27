package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/karchx/realword-nx/model"
	"github.com/karchx/realword-nx/utils"
)

func (h *Handler) GetArticle(c *fiber.Ctx) error {
	slug := c.Params("slug")
	a, err := h.articleStore.GetBySlug(slug)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if a == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newArticleResponse(userIDFromToken(c), a))
}

func (h *Handler) CreateArticle(c *fiber.Ctx) error {
	var a model.Article
	req := &articleCreateRequest{}
	if err := req.bind(c, &a, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	a.AuthorId = userIDFromToken(c)
	err := h.articleStore.CreateArticle(&a)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusCreated).JSON(newArticleResponse(userIDFromToken(c), &a))
}
