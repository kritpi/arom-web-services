package models

import (	
    "time"
    pg "github.com/lib/pq"
    "github.com/google/uuid"

)

type Diary struct {
    Id uuid.UUID `json:"id" db:"Diary_Id" `
    Date time.Time `json:"date" db:"Diary_Date"`
    Mood string `json:"mood" db:"Diary_Mood"`
    Emotions pg.StringArray `json:"emotions" db:"Diary_Emotions"`
    Description string `json:"description" db:"Diary_Description"`
    Type string `json:"type" db:"Diary_Type"`
    UserID uuid.UUID `json:"user_id" db:"User_Id"`
}