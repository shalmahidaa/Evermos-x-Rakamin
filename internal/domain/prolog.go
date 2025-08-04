package domain

type LogProduk struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ProductID   uint   `json:"product_id"`
	Nama        string `json:"nama"`
	Deskripsi   string `json:"deskripsi"`
	Harga       float64 `json:"harga"`
	Gambar      string `json:"gambar"`
}

type LogProdukRepository interface {
	Create(log *LogProduk) error
	FindByID(id uint) (*LogProduk, error)
}

type LogProdukUsecase interface {
	CreateFromProduct(p Product) (*LogProduk, error)
}