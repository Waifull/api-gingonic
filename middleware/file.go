package middleware

import (
	"gin-gonic-gorm/utils"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {

	fileHeader, _ := ctx.FormFile("file")
	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file is required.",
		})
		return
	}

	//validation by file extension
	fileExtension := []string{".jpg", ".jpg", ".jpeg", ".pdf", ".png"}
	isFileValidated := utils.FileValidationByExtension(fileHeader, fileExtension)
	if !isFileValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file not allowed.",
		})
		return
	}

	//validation by file content-type
	// fileType := []string{"image/jpg"}
	// isFileValidated := utils.FileValidation(fileHeader, fileType)
	// if !isFileValidated {
	// 	ctx.AbortWithStatusJSON(400, gin.H{
	// 		"message": "file not allowed.",
	// 	})
	// 	return
	// }

	extensionFile := filepath.Ext(fileHeader.Filename)
	filename := utils.RandomFileName(extensionFile)

	isSaved := utils.SaveFile(ctx, fileHeader, filename)
	if !isSaved {
		ctx.JSON(500, gin.H{
			"message": "internal server error, can't save file.",
		})
		return
	}

	ctx.Set("filename", filename)
	ctx.Next()
}