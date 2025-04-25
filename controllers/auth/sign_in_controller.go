package controllers

import (
	"api/models/response"
	repository "api/repositories/user/read"
	"api/validators"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *gin.Context) {
	var body validators.SignInInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Status:  400,
			Message: "Corpo JSON OU FORMDATA REQUERIDO",
			Data:    nil,
			Error: err.Error(),
		})
		return
	}

	if err := validators.Validate.Struct(body); err != nil {
		errors := validators.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, response.Response{
			Status:  400,
			Message: "CAMPO INVALIDO",
			Data:    nil,
			Error:   errors,
		})
		return
	}

	user, err := repository.FindUserByEmail(body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, response.Response{
			Status:  401,
			Message: "Dados invalidos",
			Data:    nil,
			Error:   err,
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Status:  401,
			Message: "Dados invalidos",
			Data:    nil,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

	toStringToken, err := token.SignedString([]byte(os.Getenv("SECRET")));

	if err != nil{
		c.JSON(http.StatusInternalServerError, response.Response{
			Status:  500,
			Message: "ERRO AO FAZER LOGIN",
			Data:    nil,
		})
		return
	}

	c.SetSameSite(http.SameSiteDefaultMode)
	c.SetCookie("Authorization", toStringToken, 3600 * 24 * 30, "", "", false, true)

	c.JSON(http.StatusCreated, response.Response{
		Status:  201,
		Message: "Loginefetuado com sucesso",
		Data:  response.TokenResponse{
			Token: toStringToken,
		},
	})
}
