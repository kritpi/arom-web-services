package usecases

import (
    "context"

    "github.com/kritpi/arom-web-services/configs"
    "github.com/kritpi/arom-web-services/domain/models"
    "github.com/kritpi/arom-web-services/domain/repositories"
    "github.com/kritpi/arom-web-services/domain/requests"
)

type DiaryUseCase interface {
    CreateDiary(ctx context.Context, req *requests.CreateDiaryRequest) (*models.Diary, error)
    GetAllDiary(ctx context.Context) ([]*models.Diary, error)
    GetDiaryByID(ctx context.Context, id string) (*models.Diary, error)
    GetDiaryByUserID(ctx context.Context, userID string) ([]*models.Diary, error)
    GetDiaryByDate(ctx context.Context, date string) (*models.Diary, error)
    UpdateDiary(ctx context.Context, req *requests.CreateDiaryRequest,id string) (error)
}

type diaryService struct {
    diaryRepo repositories.DiaryRepositories
    config    *configs.Config
}

func (d *diaryService) CreateDiary(ctx context.Context, req *requests.CreateDiaryRequest) (*models.Diary, error) {
    return d.diaryRepo.Create(ctx, req)
}

func (d *diaryService) GetAllDiary(ctx context.Context) ([]*models.Diary, error) {
    return d.diaryRepo.GetAll(ctx)
}
func (d *diaryService) GetDiaryByID(ctx context.Context, id string) (*models.Diary, error) {
    return d.diaryRepo.GetByID(ctx, id)
}

func (d *diaryService) GetDiaryByUserID(ctx context.Context, userID string) ([]*models.Diary, error) {
    return d.diaryRepo.GetByUserID(ctx, userID)
}

func (d *diaryService) GetDiaryByDate(ctx context.Context, date string) (*models.Diary, error) {
    return d.diaryRepo.GetByDate(ctx, date)
}

func (d *diaryService) UpdateDiary(ctx context.Context, req *requests.CreateDiaryRequest,date string) (error) {
    return d.diaryRepo.Update(ctx,req, date)
}

func ProvideDiaryService(diaryRepo repositories.DiaryRepositories, config *configs.Config) DiaryUseCase {
    return &diaryService{
        diaryRepo: diaryRepo,
        config:    config,
    }
}