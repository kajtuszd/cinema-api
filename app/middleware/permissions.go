package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/models/user"
	"net/http"
)

func Moderator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, exists := ctx.Get("user")
		fmt.Println(u)
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		if userObj, ok := u.(user.User); ok {
			if userObj.IsModerator {
				ctx.Set("user", userObj) // Set the user object in the context
				ctx.Next()
			} else {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Forbidden",
				})
				return
			}
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
