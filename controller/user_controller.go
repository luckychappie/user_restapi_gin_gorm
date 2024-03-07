package user_controller

import (
	"golang/test_rest_api/database"
	"golang/test_rest_api/model"
	"golang/test_rest_api/requests"
	"golang/test_rest_api/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUsers(c *gin.Context) {
	users := new([]model.User)

	err := database.DB.Table("users").Find(&users).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user := new(responses.UserResponse)
	err := database.DB.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

	}

	c.JSON(http.StatusOK, &user)

}

func CreateUser(c *gin.Context) {
	userRequest := new(requests.UserRequest)

	if err := c.ShouldBind(&userRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	userExist := new(model.User)

	if errExist := database.DB.Table("users").Where("email = ?", userRequest.Email).Find(&userExist).Error; errExist != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errExist)
		return
	}

	if userExist.Email != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User Already Exists"})
		return
	}

	user := new(model.User)
	user.Name = &userRequest.Name
	user.Address = &userRequest.Address
	user.Dob = &userRequest.Dob
	user.Email = &userRequest.Email

	if err := database.DB.Table("users").Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Creating user success"})

}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	userRequest := new(requests.UserRequest)

	user := new(model.User)

	if err := c.ShouldBind(&userRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
		return
	}

	if errExist := database.DB.Table("users").Where("id = ?", id).Find(&user).Error; errExist != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errExist)
		return
	}

	if user.Id == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User Does Not Exist"})
		return
	}

	user.Name = &userRequest.Name
	user.Address = &userRequest.Address
	user.Dob = &userRequest.Dob
	user.Email = &userRequest.Email

	userCreate := database.DB.Table("users").Where("id = ?", id).Updates(&user)

	if userCreate.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	if userCreate.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No changes found"})
		return
	}

	userResponse := responses.UserResponse{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Dob:     user.Dob,
	}

	c.JSON(http.StatusOK, userResponse)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user := new(model.User)

	if errExist := database.DB.Table("users").Where("id = ?", id).Find(&user).Error; errExist != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errExist)
		return
	}

	if user.Id == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User Does Not Exist"})
		return
	}

	if err := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(model.User{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "User delete success"})
}

func GetUserByPaginate(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	perPage := c.Query("perPage")
	if perPage == "" {
		perPage = "3"
	}
	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)

	if pageInt < 1 {
		pageInt = 1
	}

	users := new([]model.User)
	err := database.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users, "perPageInt": perPageInt, "pageInt": pageInt})
}
