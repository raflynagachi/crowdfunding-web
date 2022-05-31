package web

type UserResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Occupation    string `json:"occupation"`
	TokenRemember string `json:"token_remember"`
	ImageURL      string `json:"image_url"`
}
