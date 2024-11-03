package pg

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type EventPGRepository struct {
	db *sqlx.DB
}

func NewEventPGRepository(db *sqlx.DB) repositories.EventRepositories {
	return &EventPGRepository{
		db: db,
	}
}

func (e *EventPGRepository) Create(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error) {
	var event models.Event
	err := e.db.QueryRowContext(
		ctx,
		`INSERT INTO "EVENT" ("Event_Name", "Event__Description") VALUES ($1, $2) RETURNING "Event_Id", "Event_Name", "Event_Description"`,
		req.Name,
		req.Description,
	).Scan(&event.Id, &event.Name, &event.Description)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
