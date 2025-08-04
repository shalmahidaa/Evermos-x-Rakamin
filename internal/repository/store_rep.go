package repository

import (
	"evermos/config"
	"evermos/internal/domain"
)

type storeRepository struct{}

func NewStoreRepository() domain.StoreRepository {
	return &storeRepository{}
}

func (r *storeRepository) CreateForUser(userID uint, name string) error {
	store := domain.Store{
		UserID: userID,
		Name:   name,
	}
	return config.DB.Create(&store).Error
}

func (r *storeRepository) FindByUserID(userID uint) (*domain.Store, error) {
	var store domain.Store
	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *storeRepository) Update(store *domain.Store) error {
	return config.DB.Save(store).Error
}