package dtos

type UserRegisterRequestDto struct {
	Email          	string `json:"email" validate:"required,email"`
	Password       	string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UserLoginRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserLoginResponseDto struct {
	Email string `json:"email"`
	Token string `json:"token,omitempty"` // JWT token, optional for login responses
}