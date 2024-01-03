package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
	})

	return token.SignedString([]byte(viper.GetString("JWT_KEY")))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(viper.GetString("JWT_KEY")), nil
	})
	if err != nil {
		return 0, errors.New("could not parse error")
	}

	// check if the token is valid
	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	// check and retrieve token data
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	userId := int64(claims["userId"].(float64))

	return userId, nil
}
