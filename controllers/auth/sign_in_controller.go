package controllers

import (
	"api/models/response"
	repository "api/repositories/user/read"
	"api/validators"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


func SignIn(c *gin.Context){
	var body validators.SignInInput

	if err := validators.Validate.Struct(body); err != nil {
		errors := validators.FormatValidationError(err)
		c.JSON(http.StatusBadRequest,  response.Response{
			Status:  400,
			Message: "CAMPO INVALIDO",
			Data:    nil,
			Error: errors,
		})
		return
	}

	user, err := repository.FindUserByEmail(body.Email);

	if user.ID == 0{
		c.JSON(http.StatusBadRequest,  response.Response{
			Status:  401,
			Message: "Dados invalidos",
			Data:    nil,
			Error: err,
		})
		return
	}

  
	 if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil{
		c.JSON(http.StatusBadRequest,  response.Response{
			Status:  401,
			Message: "Dados invalidos",
			Data:    nil,
		})
		return
	 }

	 token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims(
		"foo": "bar",
		"nbf": time.Date(2025, 10, 10, 12, 0,0,0, time.UTC).Unix(),
	 ))
}  