package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jcmartins81/crud_HunCoding/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getuserbyid/:userId", controller.FindUserById)
	r.GET("/getuserbyemail/:userEmail", controller.FindUserByEmail)
	r.POST("/createuser", controller.CreateUser)
	r.PUT("/updateuser/:userId")
	r.DELETE("/deleteuSer/:userId")
}
