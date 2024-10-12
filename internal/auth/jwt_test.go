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
	godotenv.Load("../../.env")
	value, err := strconv.Atoi(os.Getenv("JWT_TIME"))
	assert.Nil(t, err)
	jwtData := auth.NewJWT[dto.OutputUser](os.Getenv("JWT_SECRET"), value, dto.OutputUser{
		Name:  "test",
		Email: "test@email.com",
		ID:    uuid.New(),
	})
	tokenString, err := jwtData.GenerateToken()
	assert.Nil(t, err)
	assert.NotEmpty(t, tokenString)
	err = jwtData.ValidateToken(tokenString)
	assert.Nil(t, err)
}
