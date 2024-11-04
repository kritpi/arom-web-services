package rest

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/domain/usecases"
)

type EventHandler interface {
	CreateEvent(c *fiber.Ctx) error
	GetAllEvent(c *fiber.Ctx) error
	GetByIDEvent(c *fiber.Ctx) error
	GetByUserIDEvent(c *fiber.Ctx) error
	UpdateEvent(c *fiber.Ctx) error
}

type eventHandler struct {
	service usecases.EventUseCase
}

// UpdateEvent implements EventHandler.
func (p *eventHandler) UpdateEvent(c *fiber.Ctx) error {
	var req requests.UpdateEventRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	id := c.Params("id")
	err := p.service.UpdateDateEvent(c.Context(), &req, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Event not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event updated",
	})
	
}

// GetAllEvent implements EventHandler.
func (p *eventHandler) GetAllEvent(c *fiber.Ctx) error {
	events, err := p.service.GetAllEvent(c.Context())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Event not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(events)
}

// GetByIDEvent implements EventHandler.
func (p *eventHandler) GetByIDEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	event, err := p.service.GetByIDEvent(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Event not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(event)

}

// GetByUserIDEvent implements EventHandler.
func (p *eventHandler) GetByUserIDEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	events, err := p.service.GetByUserIDEvent(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Event not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(events)
}

func (p *eventHandler) CreateEvent(c *fiber.Ctx) error {
	var req requests.CreateEventRequest
	log.Println("Request", req)
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	fmt.Printf("CreateEvent request: %v\n", req)
	event, err := p.service.CreateEvent(c.Context(), &req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

func NewEventHandler(service usecases.EventUseCase) EventHandler {
	return &eventHandler{
		service: service,
	}
}
