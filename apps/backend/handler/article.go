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
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
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

func (h *Handler) UpdateArticle(c *fiber.Ctx) error {
	slug := c.Params("slug")
	a, err := h.articleStore.GetUserArticleBySlug(userIDFromToken(c), slug)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	if a == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}

	req := &articleUpdateRequest{}
	req.populate(a)
	if err := req.bind(c, a, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	if err := h.articleStore.UpdateArticle(a); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newArticleResponse(userIDFromToken(c), a))
}

func (h *Handler) DeleteArticle(c *fiber.Ctx) error {
	slug := c.Params("slug")
	a, err := h.articleStore.GetUserArticleBySlug(userIDFromToken(c), slug)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if a == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}

	if err := h.articleStore.DeleteArticle(a); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(map[string]interface{}{"result": "ok"})
}

func (h *Handler) UndoDeleteArticle(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if err := h.articleStore.UndoDeleteArticle(slug); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	a, err := h.articleStore.GetUserArticleBySlug(userIDFromToken(c), slug)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if a == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}
	return c.Status(http.StatusOK).JSON(newArticleResponse(userIDFromToken(c), a))
}
