package repositories

import (
	"context"

	models "github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/requests"
)

type UserRepositories interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByUsername(ctx context.Context, req *requests.LoginRequest) (*models.User, error)
}

