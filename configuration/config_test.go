package configuration_test

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/configuration"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEnvData(t *testing.T) {
	godotenv.Load("../.env")

	envData := configuration.NewEnvData()

	assert.NotEmpty(t, envData)
	//api
	assert.NotEmpty(t, envData.Port)
	//database
	assert.NotEmpty(t, envData.DBHost)
	assert.NotEmpty(t, envData.DBPort)
	assert.NotEmpty(t, envData.DBUser)
	assert.NotEmpty(t, envData.DBPass)
	assert.NotEmpty(t, envData.DBName)
	//auth
	assert.NotEmpty(t, envData.JWTSecret)
	assert.NotEmpty(t, envData.JWTTime)

}
