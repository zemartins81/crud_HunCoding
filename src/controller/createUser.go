package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jcmartins81/crud_HunCoding/src/configuration/validation"
	"github.com/jcmartins81/crud_HunCoding/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("There are some incorrect field, error=%s", err.Error())
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
