package middlewares

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

func verifyToken(tokenString string) (bool, error) {
	secretKey := []byte("your-256-bit-secret")
	tokenString = tokenString[7:]
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, errors.New("1неверный токен")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, errors.New("2неверный токен")
	}
	if claims["id_user"] == nil {
		return false, errors.New("3неверный токен")
	}
	return true, nil
}

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(401, gin.H{"error": "66unauthorized"})
			ctx.Abort()
			return
		}
		valid, err := verifyToken(token)
		if err != nil {
			ctx.JSON(401, gin.H{"error": "3invalid token"})
			ctx.Abort()
			return
		}
		if !valid {
			ctx.JSON(401, gin.H{"error": "3invalid token"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
