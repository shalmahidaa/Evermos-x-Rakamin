package usecase

import (
	"evermos/internal/domain"
)

type storeUsecase struct {
	storeRepo domain.StoreRepository
}

func NewStoreUsecase(storeRepo domain.StoreRepository) domain.StoreUsecase {
	return &storeUsecase{storeRepo}
}

func (uc *storeUsecase) GetByUserID(userID uint) (*domain.Store, error) {
	return uc.storeRepo.FindByUserID(userID)
}

func (uc *storeUsecase) UpdateStore(userID uint, name string) error {
	store, err := uc.storeRepo.FindByUserID(userID)
	if err != nil {
		return err
	}
	store.Name = name
	return uc.storeRepo.Update(store)
}