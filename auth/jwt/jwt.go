package jwt

import "github.com/golang-jwt/jwt"

type JwtService interface {
	GenerateToken(userId int) (string, error)
	ParseToken(token string) (*jwt.Token, error)
}
