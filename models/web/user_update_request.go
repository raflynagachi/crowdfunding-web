package web

type UserUpdateRequest struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Password   string `json:"password"`
}
