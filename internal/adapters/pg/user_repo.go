package pg

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type UserPGRepository struct {
	db *sqlx.DB
}

func NewUserPGRepository(db *sqlx.DB) repositories.UserRepositories {
	return &UserPGRepository{
		db: db,
	}
}

// Create User
// Create User with additional logging for debugging
func (u *UserPGRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.ID = uuid.New()
	err := u.db.QueryRowContext(
		ctx,
		`INSERT INTO users (id, username, password) VALUES ($1, $2, $3) RETURNING id, username, password;`,
		user.ID,
		user.Username,
		user.Password,
	).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		// Log the UUID and error for debugging
		log.Printf("Error inserting user ID %v: %v", user.ID, err)
		return nil, err
	}
	return user, nil
}


// Get User by Username
func (u *UserPGRepository) GetUserByUsername(ctx context.Context, req *requests.LoginRequest) (*models.User, error) {
	var user models.User
	err := u.db.QueryRowContext(
		ctx,
		`SELECT id, username, password FROM users WHERE username = $1`,
		req.Username,
	).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
