package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/auth/sign-up", controllers.SignUp)

	return r
}