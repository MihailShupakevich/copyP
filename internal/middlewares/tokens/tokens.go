package tokens

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//
//type Token struct {
//	AccessToken  string `json:"access_token"`
//	RefreshToken string `json:"refresh_token"`
//}

//	func GenerateTokenPair(userId int) (Token, error) {
//		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//			"id":  userId,
//			"exp": time.Now().Add(time.Minute * 3).Unix(),
//		})
//
//		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//			"id":  userId,
//			"exp": time.Now().Add(time.Minute * 5).Unix(),
//		})
//
//		accessTokenString, err := accessToken.SignedString([]byte("secretKey"))
//		if err != nil {
//			return Token{}, err
//		}
//
//		refreshTokenString, err := refreshToken.SignedString([]byte("secretKey"))
//		if err != nil {
//			return Token{}, err
//		}
//
//		return Token{
//			AccessToken:  accessTokenString,
//			RefreshToken: refreshTokenString,
//		}, nil
//	}
func GenerateToken(userId int, expiration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RefreshToken(refreshTokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(refreshTokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil
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

	newAccessTokenString, err := newAccessToken.SignedString([]byte("secretKey"))
	if err != nil {
		return "", err
	}
	return newAccessTokenString, nil
}
