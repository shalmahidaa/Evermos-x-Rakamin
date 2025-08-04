package repository

import (
	"evermos/internal/domain"
	"evermos/config"
)

type transactionRepository struct{}

func NewTransactionRepository() domain.TransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) CreateTransaction(trx *domain.Transaction) error {
	return config.DB.Create(trx).Error
}

func (r *transactionRepository) CreateDetail(details []domain.DetailTransaction) error {
	return config.DB.Create(&details).Error
}

func (r *transactionRepository) FindByUserID(userID uint) ([]domain.Transaction, error) {
	var trxs []domain.Transaction
	err := config.DB.
		Preload("Details").
		Where("user_id = ?", userID).
		Find(&trxs).Error
	return trxs, err
}