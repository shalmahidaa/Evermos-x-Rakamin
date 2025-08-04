package handler

import (
	"evermos/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	usecase domain.CategoryUsecase
}

func NewCategoryHandler(usecase domain.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{usecase}
}

func (h *CategoryHandler) AddCategory(c *fiber.Ctx) error {
	var category domain.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.usecase.AddCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(category)
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	var category domain.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.usecase.UpdateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(category)
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	if err := h.usecase.DeleteCategory(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "category deleted"})
}

func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.usecase.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(categories)
}

func (h *CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	category, err := h.usecase.GetCategoryByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "category not found"})
	}
	return c.JSON(category)
}