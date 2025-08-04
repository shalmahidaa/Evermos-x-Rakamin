package domain

import "time"

type Address struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	PostalCode string   `json:"postal_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddressRepository interface {
	Add(address *Address) error
	GetAllByUser(userID uint) ([]Address, error)
	Update(address *Address) error
	Delete(addressID uint, userID uint) error
	GetByID(addressID uint) (*Address, error)
}

type AddressUsecase interface {
	AddAddress(address *Address) error
	GetAllByUser(userID uint) ([]Address, error)
	UpdateAddress(address *Address, userID uint) error
	DeleteAddress(addressID uint, userID uint) error
}