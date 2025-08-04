package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"evermos/internal/domain"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASS", ""),
		getEnv("DB_HOST", "127.0.0.1"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "evermos_db"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate semua model
	if err := DB.AutoMigrate(
		&domain.User{},
		&domain.Store{},
		&domain.Product{},
		&domain.Address{},
		&domain.Category{},
		&domain.Transaction{},
		&domain.DetailTransaction{},
		&domain.LogProduk{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}