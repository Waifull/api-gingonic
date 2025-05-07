package user_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/model"
	"gin-gonic-gorm/requests"
	"gin-gonic-gorm/responses"

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
	user := new(responses.UserResponse)

	errDb := database.DB.Table("users").Where("id = ?", id).First(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error.",
		})
		
		return
	}

	if user.ID == nil{
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

	userRequest := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userRequest); errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	user := new(model.User)
	user.Name = &userRequest.Name
	user.Address = &userRequest.Address
	user.BornDate = &userRequest.BornDate

	errDb := database.DB.Table("users").Create(&user).Error
	if errDb != nil{
		ctx.JSON(500, gin.H{
			"message": "can't create data.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data saved successfully.",
		"data": user,
	})

}

func Update(ctx *gin.Context){
	
}

func Delete(ctx *gin.Context){
	
}