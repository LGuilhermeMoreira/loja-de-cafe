package handler

import (
	"encoding/json"
	"errors"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/auth"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	_interface "github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/interface"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	output, err := c.CoffeeInterface.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, output)
}

func (c *CoffeeHandler) UpdateCoffee(ctx *gin.Context) {
	var input dto.InputUpdateCoffee
	if err := json.NewDecoder(ctx.Request.Body).Decode(&input); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": errors.New("invalid id").Error()})
		return
	}
	response, err := c.CoffeeInterface.Update(input, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
func (c *CoffeeHandler) DeleteCoffee(ctx *gin.Context) {}
func (c *CoffeeHandler) ListCoffees(ctx *gin.Context)  {}
func (c *CoffeeHandler) GetCoffee(ctx *gin.Context)    {}
