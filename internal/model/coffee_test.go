package model_test

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCoffee(t *testing.T) {
	coffee := model.NewCoffee(123.0, "foto do cafe", "./SomePath/SomeFile.jpg", "test 1")
	assert.NotNil(t, coffee)
	assert.NotEmpty(t, coffee)
	assert.NotEmpty(t, coffee.ID)
	assert.NotEmpty(t, coffee.Price)
	assert.NotEmpty(t, coffee.Description)
	assert.NotEmpty(t, coffee.ImagePath)
	assert.NotEmpty(t, coffee.Name)
}
