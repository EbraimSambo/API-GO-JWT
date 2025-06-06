package routes

import (
	controllers "api/controllers/auth"
	"api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/auth/sign-up", controllers.SignUp)
	r.POST("/auth/sign-in", controllers.SignIn)
	r.GET("/test", middlewares.RequireAUth,controllers.Validate)

	return r
}