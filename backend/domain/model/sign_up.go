package model

type SignupRequest struct {
	Email           string `form:"email"    binding:"required,email"`
	Password        string `form:"password" binding:"required"`
	PasswordConfirm string `form:"passwordConfirm" binding:"required"`
}

type SignupConfirmRequest struct {
	Email          string
	HashedPassword string
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
