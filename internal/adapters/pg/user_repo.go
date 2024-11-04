package pg

import (
	"context"
	"log"
	"bytes"
  "fmt"
  "io"
  "mime/multipart"
  "net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/configs"
)

type UserPGRepository struct {
	db *sqlx.DB
}

func NewUserPGRepository(db *sqlx.DB) repositories.UserRepositories {
	return &UserPGRepository{
		db: db,
	}
}

func UploadImageToSupabase(file multipart.File, fileName string, bucket string, config *configs.Config) (string, error) {
	// Read file into bytes
	fileBytes, err := io.ReadAll(file)
	if err != nil {
			return "", err
	}

	// Build Supabase Storage URL
	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", config.SUPABASE_URL, bucket, fileName)

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewReader(fileBytes))
	if err != nil {
			return "", err
	}

	req.Header.Set("Authorization", "Bearer "+config.SUPABASE_API_KEY)
	req.Header.Set("Content-Type", "image/jpeg") // Change to your specific content type if needed

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
			body, _ := io.ReadAll(resp.Body)
			return "", fmt.Errorf("upload failed: %s", body)
	}

	return fmt.Sprintf("%s/storage/v1/object/public/%s/%s", config.SUPABASE_URL, bucket, fileName), nil
}

// Create User
// Create User with additional logging for debugging
func (u *UserPGRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.ID = uuid.New()
	err := u.db.QueryRowContext(
		ctx,
		`INSERT INTO users (id, username, password, profile_image) VALUES ($1, $2, $3, $4) RETURNING id, username, password, profile_image;`,
		user.ID,
		user.Username,
		user.Password,
		user.ProfileImage,
	).Scan(&user.ID, &user.Username, &user.Password, &user.ProfileImage)

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
		`SELECT id, username, password, profile_image FROM users WHERE username = $1`,
		req.Username,
	).Scan(&user.ID, &user.Username, &user.Password, &user.ProfileImage)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
