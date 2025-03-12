package http

import (
	"Amooz/internal/book/application"
	"Amooz/internal/book/domain"
	"github.com/gofiber/fiber/v2"
	"time"
)

// BookHandler نماینده هندلر برای درخواست‌های HTTP مربوط به کتاب‌ها است
type BookHandler struct {
	service *application.BookService
}

// NewBookHandler سازنده BookHandler
func NewBookHandler(service *application.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// CreateBookHandler برای ایجاد یک کتاب جدید
func (h *BookHandler) CreateBookHandler(c *fiber.Ctx) error {
	var req struct {
		Title       string    `json:"title"`
		AuthorID    string    `json:"author_id"`
		AuthorName  string    `json:"author_name"`
		Publisher   string    `json:"publisher"`
		PublishedAt time.Time `json:"published_at"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	author := domain.Author{
		ID:   req.AuthorID,
		Name: req.AuthorName,
	}

	book, err := h.service.CreateBook(c.Context(), req.Title, author, req.Publisher, req.PublishedAt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

// FindBookByIDHandler برای پیدا کردن یک کتاب با ID
func (h *BookHandler) FindBookByIDHandler(c *fiber.Ctx) error {
	id := c.Query("id")
	book, err := h.service.FindBookByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "book not found"})
	}
	panic("implement me")
	return c.Status(fiber.StatusOK).JSON(book)
}
