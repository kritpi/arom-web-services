package pg

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/lib/pq"
)

type DiaryPGRepository struct {
	db *sqlx.DB
}

// Create implements repositories.DiaryRepositories.
func (d *DiaryPGRepository) Create(ctx context.Context, req *requests.CreateDiaryRequest) (*models.Diary, error) {
	var diary models.Diary
	emotionsArray := "{" + strings.Join(req.Emotions, ",") + "}"
	var emotions pq.StringArray
	err := d.db.QueryRowxContext(ctx, `INSERT INTO "DIARY" (
    "Diary_Date", 
    "Diary_Emotions", 
    "Diary_Mood", 
    "Diary_Description", 
    "User_Id"
) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING 
    "Diary_Id", 
    "Diary_Date", 
    "Diary_Emotions", 
    "Diary_Mood", 
    "Diary_Description", 
    "User_Id";
`,
	req.Date,emotionsArray,req.Mood,req.Description,req.UserID).Scan(&diary.Id, &diary.Date, &emotions , &diary.Mood, &diary.Description, &diary.UserID)
	if err != nil {
		return nil, err
	}
	diary.Emotions = emotions
	return &diary, nil

}
func (d* DiaryPGRepository) GetAll(ctx context.Context) ([]*models.Diary, error) {
	var diaries []*models.Diary
	err := d.db.SelectContext(ctx, &diaries, `SELECT * FROM "DIARY"`)
	if err != nil {
		return nil, err
	}
	return diaries, nil
}
// GetByDate implements repositories.DiaryRepositories.
func (d *DiaryPGRepository) GetByDate(ctx context.Context, date string) (*models.Diary, error) {
	var diary models.Diary
	err := d.db.GetContext(ctx, &diary, `SELECT "Diary_Id", "Diary_Date", "Diary_Emotions", "Diary_Mood", "Diary_Description", "User_Id" FROM "DIARY" WHERE "Diary_Date" = $1`, date)
	if err != nil {
		return nil, err
	}
	return &diary, nil
}

// GetByID implements repositories.DiaryRepositories.
func (d *DiaryPGRepository) GetByID(ctx context.Context, id string) (*models.Diary, error) {
	var diary models.Diary
	err := d.db.GetContext(ctx, &diary, `SELECT "Diary_Id", "Diary_Date", "Diary_Emotions", "Diary_Mood", "Diary_Description", "User_Id" FROM "DIARY" WHERE "Diary_Id" = $1`, id)
	if err != nil {
		return nil, err
	}
	return &diary, nil
}

// GetByUserID implements repositories.DiaryRepositories.
func (d *DiaryPGRepository) GetByUserID(ctx context.Context, userID string) ([]*models.Diary, error) {
	var diaries []*models.Diary
	err := d.db.SelectContext(ctx, &diaries, `SELECT "Diary_Id", "Diary_Date", "Diary_Emotions", "Diary_Mood", "Diary_Description", "User_Id" FROM "DIARY" WHERE "User_Id" = $1`, userID)
	if err != nil {
		return nil, err
	}
	return diaries, nil
}

// Update implements repositories.DiaryRepositories.
func (d *DiaryPGRepository) Update(ctx context.Context, req *requests.CreateDiaryRequest, date string) error {
	emotionsArrayU := "{" + strings.Join(req.Emotions, ",") + "}"
	_, err := d.db.ExecContext(ctx, `UPDATE "DIARY" SET "Diary_Date"=$1,"Diary_Emotions" = $2, "Diary_Mood" = $3, "Diary_Description" = $4 ,"User_Id"=$5 WHERE "Diary_Date" = $6`,req.Date,emotionsArrayU, req.Mood, req.Description,req.UserID,  date)
	if err != nil {
		return err
	}
	return nil
}

// Create implements repositories.DiaryRepositories.

func NewDiaryPGRepository(db *sqlx.DB) repositories.DiaryRepositories {
	return &DiaryPGRepository{
		db: db,
	}
}
