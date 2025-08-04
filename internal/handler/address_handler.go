package handler

import (
	"evermos/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type AddressHandler struct {
	addressUsecase domain.AddressUsecase
}

func NewAddressHandler(addressUsecase domain.AddressUsecase) *AddressHandler {
	return &AddressHandler{addressUsecase: addressUsecase}
}

// Add Address
func (h *AddressHandler) AddAddress(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var input domain.Address
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	input.UserID = userID

	if err := h.addressUsecase.AddAddress(&input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "address added successfully"})
}

// Get All Addresses for User
func (h *AddressHandler) GetAllAddresses(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	addresses, err := h.addressUsecase.GetAllByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(addresses)
}

// Update Address
func (h *AddressHandler) UpdateAddress(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var input domain.Address
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.addressUsecase.UpdateAddress(&input, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "address updated successfully"})
}

// Delete Address
func (h *AddressHandler) DeleteAddress(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	addressID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid address id"})
	}

	if err := h.addressUsecase.DeleteAddress(uint(addressID), userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "address deleted successfully"})
}