package handler

import (
	"encoding/json"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/auth"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	_interface "github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/interface"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	UserInterface _interface.UserInterface
	JWTInterface  auth.JWTInterface[dto.OutputUser]
}

func NewUserHandler(userInterface _interface.UserInterface, jwtInterface auth.JWTInterface[dto.OutputUser]) *UserHandler {
	return &UserHandler{
		UserInterface: userInterface,
		JWTInterface:  jwtInterface,
	}
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	var requestBody dto.InputCreateUser
	if err := json.NewDecoder(ctx.Request.Body).Decode(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	response, err := u.UserInterface.Create(requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var requestBody dto.LoginInputDto
	err := json.NewDecoder(ctx.Request.Body).Decode(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	output, err := u.UserInterface.Login(requestBody.Email, requestBody.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	tokenString, err := u.JWTInterface.GenerateToken(*output)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized,
			gin.H{"error": err.Error()})
		return
	}
	response := dto.JWTOutputDto{
		Token: tokenString,
	}
	ctx.JSON(http.StatusOK, response)
}
