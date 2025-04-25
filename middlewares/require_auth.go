package middlewares

import (
	"api/initializers"
	"api/models"
	"api/models/response"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)



func RequireAUth(c *gin.Context){
	fmt.Println("IN middleware")
	tokenString, err := c.Cookie("Authorization") 

	if err != nil{
		c.JSON(http.StatusUnauthorized, response.Response{
			Status:  201,
			Message: "Nao autorizado",
		})
	}

	token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error)  {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			c.JSON(http.StatusUnauthorized, response.Response{
				Status:  201,
				Message: "Nao autorizado",
			})
		}

		return []byte(os.Getenv("SECRET")), nil
	}) 

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		if float64(time.Now().Unix() )> claims["exp"].(float64){
			c.JSON(http.StatusUnauthorized, response.Response{
				Status:  201,
				Message: "Nao autorizado",
			})
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID ==0 {
			c.JSON(http.StatusUnauthorized, response.Response{
				Status:  201,
				Message: "Nao autorizado",
			})
		}
		c.Set("user", user)
		c.Next()
	}else{
		c.JSON(http.StatusUnauthorized, response.Response{
			Status:  201,
			Message: "Nao autorizado",
		})
	}
	c.Next();
}