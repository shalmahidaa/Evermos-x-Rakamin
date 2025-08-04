package usecase

import (
	"evermos/internal/domain"
	"fmt"
	"time"
)

type transactionUsecase struct {
	transactionRepo domain.TransactionRepository
	productRepo domain.ProductRepository
	logProdukUC domain.LogProdukUsecase
}

func NewTransactionUsecase(
	tr domain.TransactionRepository,
	pr domain.ProductRepository,
	lr domain.LogProdukUsecase,
) domain.TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: tr,
		productRepo:     pr,
		logProdukUC:     lr,
	}
}

func (uc *transactionUsecase) Checkout(userID uint, trx *domain.Transaction, details []domain.DetailTransaction) error {
	trx.UserID = userID
	trx.CreatedAt = time.Now()
	trx.KodeInvoice = fmt.Sprintf("INV-%d-%d", userID, time.Now().Unix())

	if err := uc.transactionRepo.CreateTransaction(trx); err != nil {
		return err
	}

	for i := range details {
		product, err := uc.productRepo.FindByID(details[i].ProductID)
		if err != nil {
			return fmt.Errorf("gagal ambil produk: %w", err)
		}

		logProduk, err := uc.logProdukUC.CreateFromProduct(*product)
		if err != nil {
			return fmt.Errorf("gagal buat log produk: %w", err)
		}

		details[i].TrxID = trx.ID
		details[i].LogProdukID = logProduk.ID // Set ID log produk ke detail
		details[i].CreatedAt = time.Now()
	}

	return uc.transactionRepo.CreateDetail(details)
}

func (uc *transactionUsecase) GetAllByUser(userID uint) ([]domain.Transaction, error) {
	return uc.transactionRepo.FindByUserID(userID)
}