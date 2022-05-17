package web

type AuthIsEmailAvailableRequest struct {
	Email string `json:"email" binding:"required,email"`
}
