package responses

import "github.com/jackc/pgx/v5/pgtype"

type UserResponse struct {
	ID       pgtype.UUID `json:"id" db:"ID"`
	Username pgtype.Text `json:"username" db:"Username"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
