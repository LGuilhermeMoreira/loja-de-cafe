package handler

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	_interface "github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/interface"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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

//
//func (s *ShoppingCartHandler) UpdateShoppingCart(ctx *gin.Context) {
//	shoppingCart := uuid.MustParse(ctx.Param("user_id"))
//
//}
