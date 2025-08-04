package handler

import (
	"evermos/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	usecase domain.TransactionUsecase
}

func NewTransactionHandler(uc domain.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{usecase: uc}
}

func (h *TransactionHandler) Checkout(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	userID, ok := userIDVal.(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	var input struct {
		AlamatPengirimanID uint                     `json:"alamat_pengiriman_id"`
		MethodBayar        string                   `json:"method_bayar"`
		Details            []domain.DetailTransaction `json:"details"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	trx := &domain.Transaction{
		AlamatPengirimanID: input.AlamatPengirimanID,
		MethodBayar:        input.MethodBayar,
		HargaTotal:         0,
	}

	for _, d := range input.Details {
		trx.HargaTotal += d.HargaTotal
	}

	err := h.usecase.Checkout(userID, trx, input.Details)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Transaksi berhasil", "invoice": trx.KodeInvoice})
}

func (h *TransactionHandler) GetMyTransactions(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	userID, ok := userIDVal.(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	trxs, err := h.usecase.GetAllByUser(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(trxs)
}