package controllers

import (
	"api/models/response"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Validate(c *gin.Context){
	c.JSON(http.StatusOK, response.Response{
		Status: 200,
		Message: "I LOGGIN",
		Data: "oke",
	})
}