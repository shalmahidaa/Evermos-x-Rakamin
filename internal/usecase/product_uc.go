package usecase

import (
	"errors"
	"evermos/internal/domain"
)

type productUsecase struct {
	productRepo domain.ProductRepository
	storeRepo   domain.StoreRepository
}

func NewProductUsecase(p domain.ProductRepository, s domain.StoreRepository) domain.ProductUsecase {
	return &productUsecase{
		productRepo: p,
		storeRepo:   s,
	}
}

// Add product ke toko milik user
func (uc *productUsecase) AddProduct(userID uint, product *domain.Product) error {
	store, err := uc.storeRepo.FindByUserID(userID)
	if err != nil {
		return err
	}

	product.StoreID = store.ID
	return uc.productRepo.Create(product)
}

// Get semua produk dari toko user
func (uc *productUsecase) GetProducts(userID uint) ([]domain.Product, error) {
	store, err := uc.storeRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return uc.productRepo.FindByStoreID(store.ID)
}

// Update produk, pastikan milik user
func (uc *productUsecase) UpdateProduct(userID uint, updated *domain.Product) error {
	store, err := uc.storeRepo.FindByUserID(userID)
	if err != nil {
		return err
	}

	// Ambil produk lama
	product, err := uc.productRepo.FindByID(updated.ID)
	if err != nil {
		return err
	}

	// Cek kepemilikan
	if product.StoreID != store.ID {
		return errors.New("unauthorized: product not in your store")
	}

	// Update field yang boleh
	product.Name = updated.Name
	product.Description = updated.Description
	product.Price = updated.Price
	product.Stock = updated.Stock

	return uc.productRepo.Update(product)
}

// Hapus produk jika milik user
func (uc *productUsecase) DeleteProduct(userID uint, productID uint) error {
	store, err := uc.storeRepo.FindByUserID(userID)
	if err != nil {
		return err
	}

	product, err := uc.productRepo.FindByID(productID)
	if err != nil {
		return err
	}

	if product.StoreID != store.ID {
		return errors.New("unauthorized: cannot delete others' product")
	}

	return uc.productRepo.Delete(productID)
}