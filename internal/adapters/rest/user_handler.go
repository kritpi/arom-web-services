package rest

import (
	"context" // Import context

	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/domain/usecases"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserUsecase usecases.UserUseCase // Change to use the interface directly
}

func NewUserHandler(userUsecase usecases.UserUseCase) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req requests.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Use context
	res, err := h.UserUsecase.Register(context.Background(), &req) // Pass context
    if err == nil {
        return c.Status(fiber.StatusCreated).JSON(res)
    } else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req requests.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Use context
	res, err := h.UserUsecase.Login(context.Background(), &req) // Pass context
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
