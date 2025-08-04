package routes

import (
	"evermos/internal/handler"
	"evermos/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, authHandler *handler.AuthHandler, productHandler *handler.ProductHandler, storeHandler *handler.StoreHandler, addressHandler *handler.AddressHandler, categoryHandler *handler.CategoryHandler, transactionHandler *handler.TransactionHandler) {
	api := app.Group("/api")

	// Auth routes (tidak perlu token)
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	// Protected routes (butuh token)
	protected := api.Group("", middleware.JWTProtected())

	// Profile
	protected.Get("/me", authHandler.Profile)
	protected.Put("/me", authHandler.UpdateProfile)

	// Product routes
	protected.Post("/products", productHandler.Add)
	protected.Get("/products", productHandler.GetAll)
	protected.Put("/products", productHandler.Update)
	protected.Delete("/products/:id", productHandler.Delete)

	// Toko
	protected.Get("/stores", storeHandler.GetMyStore)
	protected.Put("/stores", storeHandler.UpdateStore)

	// Address
	protected.Post("/addresses", addressHandler.AddAddress)
	protected.Get("/addresses", addressHandler.GetAllAddresses)
	protected.Put("/addresses", addressHandler.UpdateAddress)
	protected.Delete("/addresses/:id", addressHandler.DeleteAddress)

	// Category
	protected.Post("/categories", middleware.AdminOnly(), categoryHandler.AddCategory)
	protected.Put("/categories", middleware.AdminOnly(), categoryHandler.UpdateCategory)
	protected.Delete("/categories/:id", middleware.AdminOnly(), categoryHandler.DeleteCategory)
	protected.Get("/categories", categoryHandler.GetAllCategories)
	protected.Get("/categories/:id", categoryHandler.GetCategoryByID)

	//TRX
	protected.Post("/transactions", transactionHandler.Checkout)
	protected.Get("/transactions", transactionHandler.GetMyTransactions)
}