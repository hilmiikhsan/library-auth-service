package dto

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type RegisterResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
