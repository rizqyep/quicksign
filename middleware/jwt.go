package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rizqyep/quicksign/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")

		if len(token) == 2 {
			authToken := token[1]
			authorized, err := utils.IsAuthorized(authToken)
			if authorized {
				tokenClaims, err := utils.ParseTokenData(authToken)
				if err != nil {
					c.IndentedJSON(http.StatusUnauthorized, gin.H{
						"message": err.Error(),
					})
					c.Abort()
					return
				}
				c.Set("x-user-id", tokenClaims["id"])
				c.Set("x-user-username", tokenClaims["username"])
				c.Set("x-user-email", tokenClaims["email"])
				c.Next()
				return
			}
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
	}
}
