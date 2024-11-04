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

// Updatestatus implements repositories.EventRepositories.
func (e *EventPGRepository) Updatestatus(ctx context.Context, req *requests.UpdateEventRequest, id string) error {
	_, err:= e.db.ExecContext(ctx, `UPDATE "EVENT" SET "Event_Complete" = $1 WHERE "Event_Id" = $2`, req.Completed, id)
	if err != nil {
		return err
	}
	return nil
 }

// GetByUserID implements repositories.EventRepositories.
func (e *EventPGRepository) GetByUserID(ctx context.Context, id string) ([]*models.Event, error) {
	var events []*models.Event
	err := e.db.SelectContext(ctx, &events, `SELECT * FROM "EVENT" WHERE "User_Id" = $1`, id)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// GetAll implements repositories.EventRepositories.
func (e *EventPGRepository) GetAll(ctx context.Context) ([]*models.Event, error) {
	var events []*models.Event
	err := e.db.SelectContext(ctx, &events, `SELECT * FROM "EVENT"`)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// GetByID implements repositories.EventRepositories.
func (e *EventPGRepository) GetByID(ctx context.Context, id string) (*models.Event, error) {
	var event models.Event
	err := e.db.GetContext(ctx, &event, `SELECT * WHERE "Event_Id" = $1`, id)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func NewEventPGRepository(db *sqlx.DB) repositories.EventRepositories {
	return &EventPGRepository{
		db: db,
	}
}

func (e *EventPGRepository) Create(ctx context.Context, req *requests.CreateEventRequest) (*models.Event, error) {
	var event models.Event
	err := e.db.QueryRowxContext(ctx, `INSERT INTO "EVENT" (
	"Event_Title", 
	"Event_Description", 
	"Event_Start", 
	"Event_End", 
	"Event_Tag",
	"User_Id"
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
	"Event_Id",
	"Event_Title",
	"Event_Description",
	"Event_Start",
	"Event_End",
	"Event_Type",
	"Event_Complete",
	"Event_Tag",
	"User_Id";

`,
		req.Title, req.Description, req.Start, req.End,req.Tag,req.UserId).Scan(&event.Id, &event.Title, &event.Description, &event.Start, &event.End, &event.Tag,&event.Completed,&event.Type, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
