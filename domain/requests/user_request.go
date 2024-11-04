package requests

type RegisterRequest struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	ProfileImage string `json:"profile_image" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
