package file_controller

import (
	"fmt"
	"gin-gonic-gorm/utils"
	"path/filepath"
	"time"

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


	extensionFile := filepath.Ext(fileHeader.Filename)
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		
		panic(err)
	}
	currentTime := time.Now().In(loc).Format("20060102")
	filename := fmt.Sprintf("%s-%s%s",currentTime, utils.RandomString(5), extensionFile)

	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", filename))
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