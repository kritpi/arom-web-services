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
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request data"})
	}

	// Get profile image file from form
	fileHeader, err := c.FormFile("profile_image")
	if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Profile image is required"})
	}

	// Open the file for reading
	file, err := fileHeader.Open()
	if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid image file"})
	}
	defer file.Close()

	// Call use case to register user
	user, err := h.UserUsecase.Register(c.Context(), &req, file, fileHeader.Filename)
	if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
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
