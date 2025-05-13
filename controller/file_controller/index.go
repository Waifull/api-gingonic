package file_controller

import (
	"fmt"
	"gin-gonic-gorm/constanta"
	"gin-gonic-gorm/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SendStatus(ctx *gin.Context){

	filename := ctx.MustGet("filename").(string)

	ctx.JSON(200, gin.H{
		"message": "file uploaded.",
		"file_name": filename,
	})
}

func HandlerUploadFile(ctx *gin.Context){

	claimsData := ctx.MustGet("claimsData").(jwt.MapClaims)
	fmt.Println("claimsData => email => ", claimsData["email"])

	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userId => ", userId)
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
	
	ctx.JSON(200, gin.H{
		"message": "file uploaded.",
	})
}

func HandleRemoveFile(ctx *gin.Context) {

	filename := ctx.Param("filename")

	if filename == "" {
		ctx.JSON(400, gin.H{
			"message": "file name is required",
		})
	}
	err := utils.RemoveFile(constanta.DIR_FILE + filename)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "file deleted successfully.",
	})
}