package handler

import (
	"evermos/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productUC domain.ProductUsecase
}

func NewProductHandler(pu domain.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUC: pu}
}

func (h *ProductHandler) Add(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.productUC.AddProduct(userID, &product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	products, err := h.productUC.GetProducts(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.productUC.UpdateProduct(userID, &product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "product updated"})
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	productID, err := c.ParamsInt("id")
	if err != nil || productID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product ID"})
	}

	if err := h.productUC.DeleteProduct(userID, uint(productID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "product deleted"})
}