package model

type SigninRequest struct {
	Email    string `form:"email"    binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SigninResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
