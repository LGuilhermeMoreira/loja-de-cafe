package model_test

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/model"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNewShoppingCart(t *testing.T) {
	user, err := model.NewUser("ola", "test@test", "test", false)
	assert.Nil(t, err)
	log.Println(user.Id)
	shoppingCart := model.NewShoppingCart(user.Id)
	assert.NotNil(t, shoppingCart)
}

func TestShoppingCartWithItems(t *testing.T) {
	user, err := model.NewUser("ola", "test@test", "test", false)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	coffee := model.NewCoffee(100, "test", "test", "test")
	assert.NotNil(t, coffee)
	shoppingCart := model.NewShoppingCart(user.Id)
	assert.NotNil(t, shoppingCart)
	cartItem := model.NewCartItem(shoppingCart.Id, coffee.Id, 2, coffee.Price)
	assert.NotNil(t, cartItem)
	shoppingCart.CartItems = append(shoppingCart.CartItems, *cartItem)
	assert.NotEmpty(t, shoppingCart.CartItems)
}
