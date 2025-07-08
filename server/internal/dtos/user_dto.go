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

type UserUpdateRequestDto struct {
	Name     		string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Password 		string `json:"password" validate:"omitempty,min=8"`
	ClearImage      bool   `json:"clear_image,omitempty"`      // Flag to indicate if the profile picture should be cleared
}

type UserUpdateResponseDto struct {
	Name            string `json:"name"`
	ProfilePicture  string `json:"profile_picture,omitempty"` // URL or path to the profile picture
}

type UserResponseDto struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	ProfilePicture  string `json:"profile_picture,omitempty"` // URL or path to the profile picture
}