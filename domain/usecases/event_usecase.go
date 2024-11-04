package usecases

import (
	"context"

	"github.com/kritpi/arom-web-services/configs"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type EventUseCase interface {
	CreateEvent(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error)
	GetAllEvent(ctx context.Context) ([]*models.Event, error)
	GetByIDEvent(ctx context.Context, id string) (*models.Event, error)
	GetByUserIDEvent(ctx context.Context, id string) ([]*models.Event, error)
	UpdateDateEvent(ctx context.Context, req *requests.UpdateEventRequest, id string) error
}

type eventService struct {
	eventRepo repositories.EventRepositories
	config    *configs.Config
}

// UpdateDateEvent implements EventUseCase.
func (e *eventService) UpdateDateEvent(ctx context.Context, req *requests.UpdateEventRequest, id string) error {
	return e.eventRepo.Updatestatus(ctx, req, id)
}

// GetByUserIDEvent implements EventUseCase.
func (e *eventService) GetByUserIDEvent(ctx context.Context, id string) ([]*models.Event, error) {
	return e.eventRepo.GetByUserID(ctx, id)
}

// GetAll implements EventUseCase.
func (e *eventService) GetAllEvent(ctx context.Context) ([]*models.Event, error) {
	return e.eventRepo.GetAll(ctx)
}

// GetByID implements EventUseCase.
func (e *eventService) GetByIDEvent(ctx context.Context, id string) (*models.Event, error) {
	return e.eventRepo.GetByID(ctx, id)
}

// CreateEvent implements EventUseCase.
func (e *eventService) CreateEvent(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error) {
	return e.eventRepo.Create(ctx, req)
}

func ProvideEventService(eventRepo repositories.EventRepositories, config *configs.Config) EventUseCase {
	return &eventService{
		eventRepo: eventRepo,
		config:    config,
	}
}
