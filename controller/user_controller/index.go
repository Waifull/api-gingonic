package user_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/model"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {

	users := new([]model.User)
	err := database.DB.Table("users").Find(&users).Error

	if err != nil{
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error!",
		})
		return
	}
	
		ctx.JSON(200, gin.H{
			"data": users,
		})
}

func GetById(ctx *gin.Context){

	id := ctx.Param("id")
	user := new(model.User)

	errDb := database.DB.Table("users").Where("id = ?", id).First(&user).Error
	if errDb != nil || user.ID == nil{
		ctx.JSON(404, gin.H{
			"message": "data not found.",
		})
		
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data transmitted.",
		"data": user,

	})
}

func Store(ctx *gin.Context){
	
}

func Update(ctx *gin.Context){
	
}

func Delete(ctx *gin.Context){
	
}