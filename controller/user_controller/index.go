package user_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/model"
	"gin-gonic-gorm/requests"
	"gin-gonic-gorm/responses"
	"strconv"

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

	userEmailExist := new(model.User)
	database.DB.Table("users").Where("email = ?", userRequest.Email).First(&userEmailExist)
	

	if userEmailExist.Email != nil{
		ctx.JSON(400, gin.H{
			"message": "email already used.",
		})

		return
	}

	user := new(model.User)
	user.Name = &userRequest.Name
	user.Email = &userRequest.Email
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
	
	id := ctx.Param("id")
	user := new(model.User)
	userReq := new(requests.UserRequest)
	userEmailExist := new(model.User)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil{
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})

		return
	}

	errDB := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	
	if errDB != nil{
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if user.ID == nil{
		ctx.JSON(404, gin.H{
			"message": "data not found.",
		})
		return
	}

	// email exist
	errUserEmailExist := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist).Error
	if errUserEmailExist != nil{
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error.",
		})
		return
	}

	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID{

		ctx.JSON(400, gin.H{
			"message": "email already used.",
		})
		return
	}

	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil{
		ctx.JSON(500, gin.H{
			"message": "can't update data.",
		})
		return
	}

	userResponse := responses.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
	}
	ctx.JSON(200, gin.H{
		"message": "data updated successfully.",
		"data": userResponse,
	})
}

func Delete(ctx *gin.Context){
	
	id := ctx.Param("id")
	user := new(model.User)

	errFind := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errFind != nil{
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

	errDb := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&model.User{}).Error
	if errDb != nil{
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error.",
			"error": errDb.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data deleted successfully.",
	})
}

func GetUserPaginate(ctx *gin.Context){
	
	page := ctx.Query("page")
	if page == ""{
		page = "1"
	}

	perPage := ctx.Query("perPage")
	if perPage == ""{
		perPage = "10"
	}

	perPageInt, _ := strconv.Atoi(perPage)
	pageInt, _ := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	users := new([]model.User)
	err := database.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error
	if err != nil{
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": users,
		"page": pageInt,
		"per_page": perPageInt,
	})
}