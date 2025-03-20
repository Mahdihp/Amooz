package router

import (
	"Amooz/internal/user/application"
	"Amooz/internal/user/delivery/http"
	"Amooz/internal/user/infrastructure"
	"Amooz/pkg/config"
	"Amooz/pkg/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupService(cfg config.Config, app *fiber.App) {
	// ایجاد مخزن کتاب و سرویس
	userRepository := infrastructure.NewUserRepository(cfg)
	userService := application.NewUserService(userRepository)

	// ایجاد هندلرهای HTTP
	userHandler := http.NewUserHandler(userService)

	// تنظیم روت‌ها
	setupRoutes(app, userHandler, cfg.AppServer)
}

func setupRoutes(app *fiber.App, handler *http.UserHandler, cfg config.AppServer) {
	// Middleware
	api := app.Group("/api", logger.New())

	api.Get("/users", handler.GetAll)
	api.Get("/user", shared.Protected(cfg), handler.FindBookByIDHandler)

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
