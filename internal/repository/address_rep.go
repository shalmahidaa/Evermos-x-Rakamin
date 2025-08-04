package repository

import (
	"evermos/config"
	"evermos/internal/domain"
)

type addressRepository struct {}

func NewAddressRepository() domain.AddressRepository {
	return &addressRepository{}
}

func (r *addressRepository) Add(address *domain.Address) error {
	return config.DB.Create(address).Error
}

func (r *addressRepository) GetAllByUser(userID uint) ([]domain.Address, error) {
	var addresses []domain.Address
	err := config.DB.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

func (r *addressRepository) Update(address *domain.Address) error {
	return config.DB.Model(&domain.Address{}).Where("id = ? AND user_id = ?", address.ID, address.UserID).Updates(address).Error
}

func (r *addressRepository) Delete(addressID uint, userID uint) error {
	return config.DB.Where("id = ? AND user_id = ?", addressID, userID).Delete(&domain.Address{}).Error
}

func (r *addressRepository) GetByID(addressID uint) (*domain.Address, error) {
	var address domain.Address
	err := config.DB.First(&address, addressID).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}