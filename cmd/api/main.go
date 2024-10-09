package main

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/configuration"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/connection"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	envData := configuration.NewEnvData()

	_, err := connection.NewConnection(envData)

	if err != nil {
		panic(err)
	}
}
