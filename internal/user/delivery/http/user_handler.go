package http

import (
	"Amooz/internal/user/application"
	"github.com/gofiber/fiber/v2"
	"time"
)

// UserHandler نماینده هندلر برای درخواست‌های HTTP مربوط به کتاب‌ها است
type UserHandler struct {
	service application.UserRepository
}

// NewUserHandler سازنده UserHandler
func NewUserHandler(service application.UserRepository) *UserHandler {
	return &UserHandler{service: service}
}

// CreateBookHandler برای ایجاد یک کتاب جدید
func (h *UserHandler) CreateBookHandler(c *fiber.Ctx) error {
	var req struct {
		Title       string    `json:"title"`
		RoleID      int8      `json:"role_id"`
		AuthorName  string    `json:"author_name"`
		Publisher   string    `json:"publisher"`
		PublishedAt time.Time `json:"published_at"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	//role := domain.Role{
	//	ID:   req.RoleID,
	//	Name: req.AuthorName,
	//}
	//h.service.Save()
	//user, err := h.service.CreateBook(c.Context(), req.Title, role, req.Publisher, req.PublishedAt)
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	//}

	return c.Status(fiber.StatusCreated).JSON("user")
}

// FindBookByIDHandler برای پیدا کردن یک کتاب با ID
func (h *UserHandler) FindBookByIDHandler(c *fiber.Ctx) error {
	//id := c.Query("id")
	//user, err := h.service.FindBookByID(c.Context(), id)
	//if err != nil {
	//	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	//}
	return c.Status(fiber.StatusOK).JSON("user")
}
