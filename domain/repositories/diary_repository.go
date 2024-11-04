package repositories

import (
    "context"

    models "github.com/kritpi/arom-web-services/domain/models"
    "github.com/kritpi/arom-web-services/domain/requests"
)

type DiaryRepositories interface {
    Create(ctx context.Context, req *requests.CreateDiaryRequest) (*models.Diary, error)
    GetAll(ctx context.Context) ([]*models.Diary, error)
    GetByID(ctx context.Context, id string) (*models.Diary, error)
    GetByUserID(ctx context.Context, userID string) ([]*models.Diary, error)
    GetByDate(ctx context.Context, date string) (*models.Diary, error)
    Update(ctx context.Context, req *requests.UpdateDiaryRequest,date  string)(error)
    

}