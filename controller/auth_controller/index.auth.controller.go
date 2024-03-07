package authcontroller

import (
	"golang/test_rest_api/database"
	"golang/test_rest_api/model"
	"golang/test_rest_api/requests"
	"golang/test_rest_api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	uloginReq := new(requests.LoginRequest)

	if err := c.ShouldBind(uloginReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
		return
	}
	user := new(model.User)
	err := database.DB.Table("users").Where("email = ?", uloginReq.Email).Find(&user).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	if user.Email == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Email or password was wrong."})
		return
	}

	if uloginReq.Password != "12345" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Email or password was wrong."})
		return
	}

	claims := jwt.MapClaims{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token, errToken := utils.GenerateToken(&claims)

	if errToken != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": errToken})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login success", "token": token})

}
