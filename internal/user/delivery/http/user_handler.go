package http

import (
	"Amooz/internal/user/application"
	"github.com/gofiber/fiber/v2"
)

// UserHandler نماینده هندلر برای درخواست‌های HTTP مربوط به کتاب‌ها است
type UserHandler struct {
	userRepository application.IUserService
}

// NewUserHandler سازنده UserHandler
func NewUserHandler(service application.IUserService) *UserHandler {
	return &UserHandler{userRepository: service}
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {

	all, err := h.userRepository.FindAll(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})

	}
	return c.Status(fiber.StatusBadRequest).JSON(all)

}

// FindBookByIDHandler برای پیدا کردن یک کتاب با ID
func (h *UserHandler) FindBookByIDHandler(c *fiber.Ctx) error {
	//id := c.Query("id")
	//user, err := h.userRepository.FindBookByID(c.Context(), id)
	//if err != nil {
	//	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	//}
	return c.Status(fiber.StatusOK).JSON("user")
}
