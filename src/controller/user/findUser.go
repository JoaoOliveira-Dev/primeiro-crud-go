package controller

import (
	"github.com/JoaoOliveira-Dev/primeiro-crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou de forma errada")

	c.JSON(err.Code, err)
}