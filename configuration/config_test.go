package configuration_test

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/configuration"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNewEnvData(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.Nil(t, err)
	envData := configuration.NewEnvData()
	log.Println(envData)
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
	assert.IsType(t, true, envData.InProduction)
}
