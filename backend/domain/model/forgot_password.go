package model

type ForgotPasswordRequest struct {
	Email string `form:"email" binding:"required"`
}

type ForgotPasswordConfirmRequest struct {
	NewPassword        string `form:"newPassword" binding:"required"`
	NewPasswordConfirm string `form:"newPasswordConfirm" binding:"required"`
}
