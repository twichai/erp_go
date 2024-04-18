package handler

import (
	"erp/models"
	"erp/service"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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

func (h *UserHandler) LoginHandler(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	if err := h.UserService.Login(user); err != nil {
		return err
	}
	claims := jwt.MapClaims{
		"user": user,
		"role": "admin",
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY_JWT")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func (h *UserHandler) GetUserHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	user, err := h.UserService.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(user)
}
