package jwt

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"is_admin": isAdmin,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}