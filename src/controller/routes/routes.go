package routes

import (
	controller "github.com/JoaoOliveira-Dev/primeiro-crud-go/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getusers", controller.FindUser)

	r.POST("/createUser", controller.CreateUser)

	r.POST("/loginUser", controller.FindLogin)

	r.PUT("/updateusers/:id", controller.UpdateUser)

	r.DELETE("/deleteuser/:id", controller.DeleteUser)
}