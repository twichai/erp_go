package handler

import (
	"erp/models"
	"erp/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService service.UserService
}

func (h *UserHandler) CreateUserHandler(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	if err := h.UserService.CreateUser(user); err != nil {
		return err
	}
	return c.JSON(user)
}
