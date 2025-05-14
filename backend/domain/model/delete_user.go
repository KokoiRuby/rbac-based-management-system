package model

type DeleteUserRequest struct {
	UserID uint `json:"userID" form:"userID" uri:"userID"`
}

type DeleteUsersRequest struct {
	UserIDs []uint `json:"userIDs"`
}
