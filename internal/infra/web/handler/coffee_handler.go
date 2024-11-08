package handler

import (
	"encoding/json"
	"errors"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	_interface "github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/interface"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

const path = "./images/"

type CoffeeHandler struct {
	CoffeeInterface _interface.CoffeeInterface
}

func NewCoffeeHandler(coffeeInterface _interface.CoffeeInterface) *CoffeeHandler {
	return &CoffeeHandler{
		CoffeeInterface: coffeeInterface,
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
func (c *CoffeeHandler) DeleteCoffee(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": errors.New("invalid id").Error()})
		return
	}
	if err := c.CoffeeInterface.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": errors.New("invalid id").Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)

}
func (c *CoffeeHandler) ListCoffees(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		limit = 5
	}
	sort := ctx.Param("sort")

	response, err := c.CoffeeInterface.FindAll(page, limit, sort)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
func (c *CoffeeHandler) GetCoffee(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := c.CoffeeInterface.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
