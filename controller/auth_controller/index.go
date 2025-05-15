package auth_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/model"
	"gin-gonic-gorm/requests"
	"gin-gonic-gorm/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *gin.Context){

	loginReq := new(requests.LoginRequest)
	if errReq := ctx.ShouldBind(&loginReq); errReq != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	
	user := new(model.User)
	errUser := database.DB.Table("users").Where("email = ?", loginReq.Email).Find(&user).Error
	if errUser != nil {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "credential not valid.",
		})
		return
	}

	// check password
	if loginReq.Password != "12345" {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "credential not valid.",
		})
		return
	}

	claims := jwt.MapClaims{
		"id": user.ID,
		"name": user.Name,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token, errToken := utils.GenerateToken(&claims)
	if errToken != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "failed generate token.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Login successfully.",
		"token": token,
	})

}