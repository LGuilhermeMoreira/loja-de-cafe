package handler

import (
	"encoding/json"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/auth"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	_interface "github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/interface"
	"github.com/LGuilhermeMoreira/loja-de-cafe/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const path = "./images/"

type CoffeeHandler struct {
	CoffeeInterface _interface.CoffeeInterface
	JWTInterface    auth.JWTInterface[dto.OutputUser]
}

func NewCoffeeHandler(coffeeInterface _interface.CoffeeInterface, jwtInterface auth.JWTInterface[dto.OutputUser]) *CoffeeHandler {
	return &CoffeeHandler{
		CoffeeInterface: coffeeInterface,
		JWTInterface:    jwtInterface,
	}
}

func (c *CoffeeHandler) CreateCoffee(ctx *gin.Context) {
	var input dto.InputCreateCoffee
	if err := json.NewDecoder(ctx.Request.Body).Decode(&input); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	pathFile, err := utils.SaveImage(input.Data, path, input.LabelFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	output, err := c.CoffeeInterface.Create(input, pathFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, output)
}
