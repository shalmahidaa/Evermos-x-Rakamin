package usecase

import (
	"evermos/internal/domain"
	"fmt"
)

type addressUsecase struct {
	addressRepo domain.AddressRepository
}

func NewAddressUsecase(addressRepo domain.AddressRepository) domain.AddressUsecase {
	return &addressUsecase{addressRepo: addressRepo}
}

func (uc *addressUsecase) AddAddress(address *domain.Address) error {
	return uc.addressRepo.Add(address)
}

func (uc *addressUsecase) GetAllByUser(userID uint) ([]domain.Address, error) {
	return uc.addressRepo.GetAllByUser(userID)
}

func (uc *addressUsecase) UpdateAddress(address *domain.Address, userID uint) error {
	// Cek apakah address tersebut milik user
	addr, err := uc.addressRepo.GetByID(address.ID)
	if err != nil {
		return err
	}

	if addr.UserID != userID {
		return fmt.Errorf("unauthorized: address does not belong to user")
	}

	// Lanjut update
	return uc.addressRepo.Update(address)
}

func (uc *addressUsecase) DeleteAddress(addressID uint, userID uint) error {
	return uc.addressRepo.Delete(addressID, userID)
}