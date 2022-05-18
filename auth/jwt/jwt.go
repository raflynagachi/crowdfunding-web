package jwt

type JwtService interface {
	GenerateToken(userId int) (string, error)
	ParseToken(token string) (string, error)
}
