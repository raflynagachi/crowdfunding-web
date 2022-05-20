package jwt

import (
	"errors"

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

func (service *JwtServiceImpl) ParseToken(token string) (*jwt.Token, error) {
	jwtSecretKey := config.GetSecret()
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return jwtToken, err
	}
	return jwtToken, nil
}

func (service *JwtServiceImpl) JwtTokenToMap(token *jwt.Token) (jwt.MapClaims, error) {
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return claim, errors.New("claim doesn't exist")
	}
	return claim, nil
}
