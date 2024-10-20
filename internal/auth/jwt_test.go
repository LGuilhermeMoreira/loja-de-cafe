package auth_test

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/auth"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func TestJWT(t *testing.T) {
	err := godotenv.Load("../../.env")
	assert.Nil(t, err)
	value, err := strconv.Atoi(os.Getenv("JWT_TIME"))
	assert.Nil(t, err)
	jwtData := auth.NewJWT[dto.OutputUser](os.Getenv("JWT_SECRET"), value)
	tokenString, err := jwtData.GenerateToken(dto.OutputUser{
		Name:  "test",
		Email: "test@email.com",
		ID:    uuid.New(),
		Admin: false,
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, tokenString)
	err = jwtData.ValidateToken(tokenString)
	assert.Nil(t, err)
}

func TestJWT_ValidateTokenAdmin(t *testing.T) {
	err := godotenv.Load("../../.env")
	assert.Nil(t, err)
	value, err := strconv.Atoi(os.Getenv("JWT_TIME"))
	assert.Nil(t, err)
	jwtData := auth.NewJWT[dto.OutputUser](os.Getenv("JWT_SECRET"), value)
	tokenString, err := jwtData.GenerateToken(dto.OutputUser{
		Name:  "test",
		Email: "test@email.com",
		ID:    uuid.New(),
		Admin: true,
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, tokenString)
	err = jwtData.ValidateTokenAdmin(tokenString)
	assert.Nil(t, err)
}
