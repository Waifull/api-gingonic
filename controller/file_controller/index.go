package file_controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandlerUploadFile(ctx *gin.Context){

	fileHeader, _ := ctx.FormFile("file")
	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file is required.",
		})
		return
	}

	// file, errFile := fileHeader.Open()
	// if errFile != nil {
	// 	panic(errFile)
	// }

	// extensionFile := filepath.Ext(fileHeader.Filename)
	// filename := ""

	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", fileHeader.Filename))
	if errUpload != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error, can't save file.",
		})
		return
	}
	
	ctx.JSON(200, gin.H{
		"message": "file uploaded.",
	})
}