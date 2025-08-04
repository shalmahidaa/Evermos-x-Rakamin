package domain

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	StoreID     uint      `json:"store_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductRepository interface {
	Create(product *Product) error
	FindByStoreID(storeID uint) ([]Product, error)
	FindByID(id uint) (*Product, error)
	Update(product *Product) error
	Delete(id uint) error
}

type ProductUsecase interface {
	AddProduct(userID uint, product *Product) error
	GetProducts(userID uint) ([]Product, error)
	UpdateProduct(userID uint, product *Product) error
	DeleteProduct(userID uint, productID uint) error
}