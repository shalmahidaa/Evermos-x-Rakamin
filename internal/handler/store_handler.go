package handler

import (
	"evermos/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type StoreHandler struct {
	storeUsecase domain.StoreUsecase
}

func NewStoreHandler(storeUsecase domain.StoreUsecase) *StoreHandler {
	return &StoreHandler{storeUsecase}
}

func (h *StoreHandler) GetMyStore(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	store, err := h.storeUsecase.GetByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "store not found"})
	}

	return c.JSON(store)
}

type UpdateStoreRequest struct {
	Name string `json:"name"`
}

func (h *StoreHandler) UpdateStore(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req UpdateStoreRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.storeUsecase.UpdateStore(userID, req.Name); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update store"})
	}

	return c.JSON(fiber.Map{"message": "store updated successfully"})
}