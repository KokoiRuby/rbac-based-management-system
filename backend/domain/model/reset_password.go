package model

type ResetPasswordRequest struct {
	OldPassword string `form:"oldPassword" binding:"required"`
	NewPassword string `form:"newPassword" binding:"required"`
}
