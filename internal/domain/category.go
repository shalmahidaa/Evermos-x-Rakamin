package domain

import "time"

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryRepository interface {
	Create(category *Category) error
	Update(category *Category) error
	Delete(id uint) error
	GetAll() ([]Category, error)
	GetByID(id uint) (*Category, error)
}

type CategoryUsecase interface {
	AddCategory(category *Category) error
	UpdateCategory(category *Category) error
	DeleteCategory(id uint) error
	GetAllCategories() ([]Category, error)
	GetCategoryByID(id uint) (*Category, error)
}