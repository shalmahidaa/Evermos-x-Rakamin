package handler

import (
	"evermos/internal/domain"
	//"github.com/gofiber/fiber/v2"
)

type LogProdukHandler struct {
	usecase domain.LogProdukUsecase
}

func NewLogProdukHandler(uc domain.LogProdukUsecase) *LogProdukHandler {
	return &LogProdukHandler{usecase: uc}
}