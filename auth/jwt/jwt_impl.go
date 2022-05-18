package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/raflynagachi/crowdfunding-web/app/config"
)

type JwtServiceImpl struct {
}

func NewJwtService() JwtService {
	return &JwtServiceImpl{}
}

func (service *JwtServiceImpl) GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecretKey := config.GetSecret()
	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (service *JwtServiceImpl) ParseToken(token string) (string, error) {
	panic("not implemented") // TODO: Implement
}
