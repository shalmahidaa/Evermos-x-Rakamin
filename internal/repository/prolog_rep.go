package repository

import (
	"evermos/config"
	"evermos/internal/domain"
)

type logProdukRepository struct{}

func NewLogProdukRepository() domain.LogProdukRepository {
	return &logProdukRepository{}
}

func (r *logProdukRepository) Create(log *domain.LogProduk) error {
	return config.DB.Create(log).Error
}

func (r *logProdukRepository) FindByID(id uint) (*domain.LogProduk, error) {
	var log domain.LogProduk
	err := config.DB.First(&log, id).Error
	return &log, err
}