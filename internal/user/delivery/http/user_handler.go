package http

import (
	"Amooz/internal/user/application"
	"github.com/gofiber/fiber/v2"
)

// UserHandler نماینده هندلر برای درخواست‌های HTTP مربوط به کتاب‌ها است
type UserHandler struct {
	userService application.IUserService
}

func NewUserHandler(service application.IUserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {

	all, err := h.userService.FindAll(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})

	}
	return c.Status(fiber.StatusBadRequest).JSON(all)

}

// FindBookByIDHandler برای پیدا کردن یک کتاب با ID
func (h *UserHandler) FindBookByIDHandler(c *fiber.Ctx) error {
	//id := c.Query("id")
	//user, err := h.userService.FindBookByID(c.Context(), id)
	//if err != nil {
	//	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	//}
	return c.Status(fiber.StatusOK).JSON("user")
}
