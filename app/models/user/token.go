package user

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var GenerateTokenError = errors.New("failed to create token")

func GenerateToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return tokenStr, GenerateTokenError
	}
	return tokenStr, nil
}
