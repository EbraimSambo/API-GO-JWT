package controllers

import (
	"api/models/response"
	repository "api/repositories/user/read"
	services "api/services/user/register"
	"api/validators"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	var body validators.SignUpInput

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
		c.JSON(http.StatusBadRequest,  response.Response{
			Status:  400,
			Message: "CAMPO INVALIDO",
			Data:    nil,
			Error: errors,
		})
		return
	}
	
	if err := validators.Validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Status:  400,
			Message: "CAMPO INVALIDO",
			Data:    nil,
			Error: err.Error(),
		})
		return
	}

	if existingUser, _ := repository.FindUserByEmail(body.Email); existingUser != nil {
		c.JSON(http.StatusConflict, response.Response{
			Status:  409,
			Message: "Ja existe um usuario com este email",
			Data:    nil,
			Error: "Duplicacao de email",
		})
		return
	}
	
	user, err := services.RegisterUser(body.Name, body.Email, body.Password)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			c.JSON(http.StatusConflict, response.Response{
				Status:  409,
				Message: "Ja existe um usuario com este email",
				Data:    nil,
				Error: err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, response.Response{
				Status:  500,
				Message: "Erro ao criar usuario",
				Data:    nil,
				Error: err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  201,
		Message: "Usuário criado com sucesso",
		Data:  response.ResponseUser{
			Id: int(user.ID),
			Name: user.Name,
			Email: user.Name,
		},
	})
}
