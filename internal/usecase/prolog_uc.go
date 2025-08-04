package usecase

import (
	"evermos/internal/domain"
)

type logProdukUsecase struct {
	repo domain.LogProdukRepository
}

func NewLogProdukUsecase(r domain.LogProdukRepository) domain.LogProdukUsecase {
	return &logProdukUsecase{repo: r}
}

func (uc *logProdukUsecase) CreateFromProduct(p domain.Product) (*domain.LogProduk, error) {
	log := &domain.LogProduk{
		ProductID:  p.ID,
		Nama:       p.Name,
		Deskripsi:  p.Description,
		Harga:      p.Price,
		Gambar:     "", // jika kamu punya field gambar, ambil dari produk
	}
	err := uc.repo.Create(log)
	return log, err
}