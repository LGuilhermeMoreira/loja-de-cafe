package connection_test

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/configuration"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/connection"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConnection(t *testing.T) {
	godotenv.Load("../../../../.env")
	envData := configuration.NewEnvData()
	assert.NotEmpty(t, envData)
	db, err := connection.NewConnection(envData)
	assert.NoError(t, err)
	assert.NotNil(t, db)
}
