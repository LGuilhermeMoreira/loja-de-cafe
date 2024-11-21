package handler

import (
	"encoding/json"
	"net/http"

	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	_interface "github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/interface"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ShoppingCartHandler struct {
	ShoppingCartInterface _interface.ShoppingCartInterface
}

func (s *ShoppingCartHandler) CreateShoppingCart(ctx *gin.Context) {
	userID := uuid.MustParse(ctx.Param("user_id"))
	shoppingCartID, err := s.ShoppingCartInterface.Create(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, dto.OutputCreateShoppingCartDto{
		Id: shoppingCartID.String(),
	})
}

func (s *ShoppingCartHandler) UpdateShoppingCart(ctx *gin.Context) {
	shoppingCartID := uuid.MustParse(ctx.Param("shopping_cart_id"))
	var coffeeDto dto.InputAddItemShoppingCartDto
	err := json.NewDecoder(ctx.Request.Body).Decode(&coffeeDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := s.ShoppingCartInterface.AddItem(shoppingCartID, coffeeDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}
