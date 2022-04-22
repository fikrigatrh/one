package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"number1/auth"
	"number1/models"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ResponseErrorCustom{
				ResponseCode:    "401",
				ResponseMessage: "unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
