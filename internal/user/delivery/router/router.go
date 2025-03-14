package router

import (
	"Amooz/internal/user/application"
	"Amooz/internal/user/delivery/http"
	"Amooz/internal/user/infrastructure"
	"Amooz/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupService(cfg config.Config, app *fiber.App) {
	// ایجاد مخزن کتاب و سرویس
	userRepository := infrastructure.NewUserRepository(cfg)
	userService := application.NewBookService(userRepository)

	// ایجاد هندلرهای HTTP
	userHandler := http.NewUserHandler(userService)

	// تنظیم روت‌ها
	setupRoutes(app, userHandler)
}

func setupRoutes(app *fiber.App, handler *http.UserHandler) {
	// Middleware
	api := app.Group("/api", logger.New())

	api.Post("/users", handler.CreateBookHandler) // ایجاد کتاب
	api.Get("/user", handler.FindBookByIDHandler) // پیدا کردن کتاب با ID

	//api.Get("/", handler.Hello)

	// Auth
	//auth := api.Group("/auth")
	//auth.Post("/login", handler.Login)

	// User
	//user := api.Group("/user")
	//user.Get("/:id", handler.GetUser)
	//user.Post("/", handler.CreateUser)
	//user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	//user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Product
	//product := api.Group("/product")
	//product.Get("/", handler.GetAllProducts)
	//product.Get("/:id", handler.GetProduct)
	//product.Post("/", middleware.Protected(), handler.CreateProduct)
	//product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
