package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id          uuid.UUID `json:"id" db:"Event_Id"`
	Title       string    `json:"title" db:"Event_Title"`
	Description string    `json:"description" db:"Event_Description"`
	Start       time.Time `json:"start" db:"Event_Start"`
	End         time.Time `json:"end" db:"Event_End"`
	Type        string    `json:"type" db:"Event_Type"`
	Tag         string    `json:"tag" db:"Event_Tag"`
	Completed   bool      `json:"completed" db:"Event_Complete"`
	UserId      uuid.UUID `json:"user_id" db:"User_Id"`
}
