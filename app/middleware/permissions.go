package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/models/ticket"
	"github.com/kajtuszd/cinema-api/app/models/user"
	"net/http"
)

func abortWithStatusForbidden(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "Forbidden",
	})
}

func Moderator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, exists := ctx.Get("user")
		if !exists {
			abortWithStatusForbidden(ctx)
			return
		}
		if u, ok := u.(user.User); ok && u.IsModerator {
			ctx.Next()
		} else {
			abortWithStatusForbidden(ctx)
			return
		}
	}
}

func AccountOwnerOrModerator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, exists := ctx.Get("user")
		if !exists {
			abortWithStatusForbidden(ctx)
			return
		}
		if us, ok := u.(user.User); ok {
			username := ctx.Param("username")
			if us.Username == username || us.IsModerator {
				ctx.Next()
				return
			}
		}
		abortWithStatusForbidden(ctx)
		return
	}
}

func TicketOwnerOrModerator(ticketService ticket.TicketService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, exists := ctx.Get("user")
		if !exists {
			abortWithStatusForbidden(ctx)
			return
		}
		if u, ok := u.(user.User); ok {
			ticketID := ctx.Param("id")
			t, err := ticketService.GetByID(ticketID)
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			if t.Owner.Username == u.Username || u.IsModerator {
				ctx.Next()
				return
			}
		}
		abortWithStatusForbidden(ctx)
		return
	}
}
