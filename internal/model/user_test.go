package model_test

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/model"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := model.NewUser("a",
		"b",
		"c",
		false)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Id)
	assert.Equal(t, "a", user.Name)
	assert.Equal(t, "b", user.Email)
	assert.Equal(t, false, user.Admin)
	assert.True(t, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("c")) == nil)
}
