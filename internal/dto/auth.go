package dto

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=8,max=50"`
	FullName string `json:"full_name" validate:"required,min=2,max=50"`
}

type RegisterResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type LoginResponse struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	FullName     string `json:"full_name"`
	Role         string `json:"role"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}
