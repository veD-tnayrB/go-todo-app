package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	rate_limiter "github.com/veD-tnayrB/todo-app/common/rate-limiter"
	"github.com/veD-tnayrB/todo-app/common/responses"
)

func RateLimiterMiddleware(rateLimiter *rate_limiter.RateLimiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !rateLimiter.CheckSpace() {
			ctx.JSON(http.StatusTooManyRequests, responses.Error{Code: http.StatusTooManyRequests, Status: false})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
