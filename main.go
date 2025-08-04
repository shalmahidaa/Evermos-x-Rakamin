package main

import (
	"evermos/config"
	"evermos/internal/handler"
	"evermos/internal/repository"
	"evermos/internal/usecase"
	"evermos/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.InitDB()

	app := fiber.New()

	// Repositories
	userRepo := repository.NewUserRepository()
	storeRepo := repository.NewStoreRepository()
	addressRepo := repository.NewAddressRepository()
	categoryRepo := repository.NewCategoryRepository()
	productRepo := repository.NewProductRepository()
	logProdukRepo := repository.NewLogProdukRepository()
	transactionRepo := repository.NewTransactionRepository()

	// Usecases
	userUC := usecase.NewUserUsecase(userRepo, storeRepo)
	storeUC := usecase.NewStoreUsecase(storeRepo)
	addressUC := usecase.NewAddressUsecase(addressRepo)
	categoryUC := usecase.NewCategoryUsecase(categoryRepo)
	productUC := usecase.NewProductUsecase(productRepo, storeRepo)
	logProdukUC := usecase.NewLogProdukUsecase(logProdukRepo)
	transactionUC := usecase.NewTransactionUsecase(transactionRepo, productRepo, logProdukUC)

	// Handlers
	authHandler := handler.NewAuthHandler(userUC)
	storeHandler := handler.NewStoreHandler(storeUC)
	addressHandler := handler.NewAddressHandler(addressUC)
	categoryHandler := handler.NewCategoryHandler(categoryUC)
	productHandler := handler.NewProductHandler(productUC)
	transactionHandler := handler.NewTransactionHandler(transactionUC)

	// Routes
	routes.Setup(app, authHandler, productHandler, storeHandler, addressHandler, categoryHandler,transactionHandler)

	app.Listen(":8080")
}