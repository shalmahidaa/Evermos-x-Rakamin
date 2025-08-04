package repository

import (
	"evermos/config"
	"evermos/internal/domain"
	//"gorm.io/gorm"
)

type categoryRepository struct {}

func NewCategoryRepository() domain.CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) Create(category *domain.Category) error {
	return config.DB.Create(category).Error
}

func (r *categoryRepository) Update(category *domain.Category) error {
	return config.DB.Model(&domain.Category{}).
		Where("id = ?", category.ID).
		Update("name", category.Name).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return config.DB.Delete(&domain.Category{}, id).Error
}

func (r *categoryRepository) GetAll() ([]domain.Category, error) {
	var categories []domain.Category
	err := config.DB.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetByID(id uint) (*domain.Category, error) {
	var category domain.Category
	err := config.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}