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

	assert.Equal(t, envData.Port, "3000")
	assert.Equal(t, envData.DBHost, "postgres")
	assert.Equal(t, envData.DBPort, "3001")
	assert.Equal(t, envData.DBUser, "postgres")
	assert.Equal(t, envData.DBPass, "password")
	assert.Equal(t, envData.DBName, "database")
	assert.Equal(t, envData.JWTSecret, "12")
	assert.Equal(t, envData.JWTTime, "123")

}
