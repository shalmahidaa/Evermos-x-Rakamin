package domain

import "time"

type Store struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StoreRepository interface {
	CreateForUser(userID uint, name string) error
	FindByUserID(userID uint) (*Store, error)
	Update(store *Store) error
}

type StoreUsecase interface {
	GetByUserID(userID uint) (*Store, error)
	UpdateStore(userID uint, name string) error
}