package routes

import (
	controllers "api/controllers/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/auth/sign-up", controllers.SignUp)

	return r
}