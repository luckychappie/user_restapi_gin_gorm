package middleware

import (
	"golang/test_rest_api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")

	if strings.Contains(bearerToken, "Bearer") {
		c.AbortWithStatusJSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "Invalid token."})
		return
	}
	token := strings.Replace(bearerToken, "Bearer ", "", -1)
	if token == "" {
		c.AbortWithStatusJSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "unauthenticated."})
		return
	}

	claimsData, err := utils.DecodeToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "Invalid token."})
		return
	}

	c.Set("claimsData", claimsData)
	c.Set("user_id", claimsData["id"])
	c.Set("user_name", claimsData["name"])
	c.Set("user_email", claimsData["email"])

	if token != "123" {
		c.AbortWithStatusJSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "Invalid token."})
		return
	}

	c.Next()
}
