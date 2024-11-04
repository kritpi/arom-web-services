package repositories

import (
	"context"

	models "github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type EventRepositories interface {
	Create(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error)
	GetAll(ctx context.Context) ([]*models.Event, error)
	GetByID(ctx context.Context, id string) (*models.Event, error)
	GetByUserID(ctx context.Context, id string) ([]*models.Event, error)
	Updatestatus(ctx context.Context,req *requests.UpdateEventRequest, id string) error
	
}
