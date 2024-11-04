package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"User_Id"`
	Username     string    `json:"username" db:"Username"`
	Password     string    `json:"password" db:"Password"`
	ProfileImage string    `json:"profile_image" db:"Profile_Image"`
}
