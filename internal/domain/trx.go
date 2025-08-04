package domain

import "time"

type Transaction struct {
	ID              uint
	UserID          uint
	AlamatPengirimanID uint
	HargaTotal      int
	KodeInvoice     string
	MethodBayar     string
	CreatedAt       time.Time

	Details []DetailTransaction `json:"details" gorm:"foreignKey:TrxID"`
}

type DetailTransaction struct {
	ID         uint
	TrxID      uint
	ProductID	uint
	LogProdukID uint
	TokoID     uint
	Kuantitas  int
	HargaTotal int
	CreatedAt  time.Time
}

type TransactionRepository interface {
	CreateTransaction(trx *Transaction) error
	CreateDetail(details []DetailTransaction) error
	FindByUserID(userID uint) ([]Transaction, error)
}

type TransactionUsecase interface {
	Checkout(userID uint, trx *Transaction, details []DetailTransaction) error
	GetAllByUser(userID uint) ([]Transaction, error)
}