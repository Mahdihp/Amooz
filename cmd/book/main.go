package main

import (
	"Amooz/internal/book/application"
	"Amooz/internal/book/delivery/http"
	"Amooz/internal/book/infrastructure"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// ایجاد مخزن کتاب و سرویس
	bookRepo := infrastructure.NewBookRepository() // استفاده از BookRepository
	bookService := application.NewBookService(bookRepo)

	// ایجاد هندلرهای HTTP
	bookHandler := http.NewBookHandler(bookService)

	// ایجاد برنامه Fiber
	app := fiber.New()

	// تنظیم روت‌ها
	app.Post("/books", bookHandler.CreateBookHandler) // ایجاد کتاب
	app.Get("/book", bookHandler.FindBookByIDHandler) // پیدا کردن کتاب با ID

	// شروع سرور
	log.Println("Starting server on :8080...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
