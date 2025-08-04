package repository

import (
	"evermos/config"
	"evermos/internal/domain"
)

type productRepository struct{}

func NewProductRepository() domain.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Create(product *domain.Product) error {
	return config.DB.Create(product).Error
}

func (r *productRepository) FindByStoreID(storeID uint) ([]domain.Product, error) {
	var products []domain.Product
	err := config.DB.Where("store_id = ?", storeID).Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := config.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *domain.Product) error {
	return config.DB.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return config.DB.Delete(&domain.Product{}, id).Error
}