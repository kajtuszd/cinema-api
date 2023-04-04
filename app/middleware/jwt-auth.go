package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kajtuszd/cinema-api/app/initializers"
	"github.com/kajtuszd/cinema-api/app/models"
	"net/http"
	"os"
	"time"
)

func JWTAuth(db *initializers.GormDatabase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			var user models.User
			db.DB().First(&user, claims["sub"])

			if user.ID == 0 {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.Set("user", user)
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
