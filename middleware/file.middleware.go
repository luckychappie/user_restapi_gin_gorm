package middleware

import (
	"fmt"
	"golang/test_rest_api/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
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
	c.Set("filename", fileName)
	c.Next()
}
