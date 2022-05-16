package web

type UserCreateRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Password   string `json:"password"`
}
