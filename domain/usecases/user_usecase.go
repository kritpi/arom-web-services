package usecases

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kritpi/arom-web-services/configs"
	"github.com/kritpi/arom-web-services/domain/models"
	"github.com/kritpi/arom-web-services/internal/adapters/pg"

	"github.com/kritpi/arom-web-services/domain/repositories"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/domain/responses"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	// Register(ctx context.Context, req *requests.RegisterRequest) (*responses.UserResponse, error)
	Register(ctx context.Context, req *requests.RegisterRequest, file multipart.File, fileName string) (*responses.UserResponse, error)
	Login(ctx context.Context, req *requests.LoginRequest) (*responses.LoginResponse, error)
}

type userService struct {
	userRepo repositories.UserRepositories
	config   *configs.Config
}

// Register implements UserUseCase.
func (u *userService) Register(ctx context.Context, req *requests.RegisterRequest, file multipart.File, fileName string) (*responses.UserResponse, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
			return nil, err
	}

	// Check if username already exists
	loginReq := &requests.LoginRequest{Username: req.Username, Password: req.Password}
	existingUser, _ := u.userRepo.GetUserByUsername(ctx, loginReq)
	if existingUser != nil {
			return nil, errors.New("username already exists")
	}

	// Initialize image URL variable
	var profileImageURL string
	if file != nil {
			// Upload image to Supabase
			profileImageURL, err = pg.UploadImageToSupabase(file, fileName, u.config.SUPABASE_BUCKET, u.config)
			if err != nil {
					return nil, err
			}
	}

	// Create new user model
	user := &models.User{
			Username:     req.Username,
			Password:     string(hashedPassword),
			ProfileImage: profileImageURL,
	}

	// Save user to repository
	if _, err := u.userRepo.CreateUser(ctx, user); err != nil {
			return nil, err
	}

	// Prepare response
	return &responses.UserResponse{
			ID:           pgtype.UUID{Bytes: [16]byte(user.ID)},
			Username:     pgtype.Text{String: user.Username, Valid: true},
			ProfileImage: pgtype.Text{String: profileImageURL, Valid: true},
	}, nil
}

// Login implements UserUseCase.
func (u *userService) Login(ctx context.Context, req *requests.LoginRequest) (*responses.LoginResponse, error) {
	// Retrieve user from repository by username
	user, err := u.userRepo.GetUserByUsername(ctx, req)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Compare provided password with stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := u.generateJWT(user.ID.String(), user.Username, user.ProfileImage)
	if err != nil {
		return nil, err
	}

	// Prepare response
	return &responses.LoginResponse{Token: token}, nil
}

func (u *userService) generateJWT(userID, username, profileImage string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"ProfileImage": profileImage,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(u.config.JWT_SECRET))
}

// ProvideUserService is a factory function to create a new UserUseCase.
func ProvideUserService(userRepo repositories.UserRepositories, config *configs.Config) UserUseCase {
	return &userService{
		userRepo: userRepo,
		config:   config,
	}
}
