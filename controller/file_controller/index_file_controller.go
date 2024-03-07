package file_controller

import (
	"fmt"
	"golang/test_rest_api/constantvar"
	"golang/test_rest_api/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HandleUploadedFile(c *gin.Context) {

	claimsData := c.MustGet("claimsData").(jwt.MapClaims)
	fmt.Println("claimsData => email => ", claimsData["email"])

	userId := c.MustGet("user_id").(int)
	fmt.Println("claimsData => ID => ", userId)

	uploadedFile, _ := c.FormFile("file")

	if uploadedFile == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No uploaded file found"})
		return
	}

	fileExt := filepath.Ext(uploadedFile.Filename)
	fileName := fmt.Sprintf("%s%s", utils.RandomString(7), fileExt)

	err := c.SaveUploadedFile(uploadedFile, fmt.Sprintf("./public/files/%s", fileName))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "File upload fail!"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded."})
}

func HandleRemoveFile(c *gin.Context) {
	filename := c.Param("filename")

	if filename == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Please fill deleted file name."})
		return
	}

	err := utils.RemoveFile(constantvar.DIR_FILE + filename)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Delete fail"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted"})
}

func SendStatus(c *gin.Context) {
	filename := c.MustGet("filename").(string)

	c.JSON(http.StatusOK, gin.H{"message": "File Uplaoded.", "file name": filename})
}
