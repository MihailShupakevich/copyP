package tokens

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(userId int, expiration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":            userId,
		"exp":           time.Now().Add(expiration).Unix(),
		"authorization": true,
	})

	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RefreshToken(refreshTokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(refreshTokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", err
	}

	userId := int(claims["id"].(float64))

	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	})

	newAccessTokenString, err := newAccessToken.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}
	return newAccessTokenString, nil
}
