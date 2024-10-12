package main

import (
	"errors"
	"github.com/LGuilhermeMoreira/loja-de-cafe/configuration"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/connection"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/migration"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	envData := configuration.NewEnvData()
	db, err := connection.NewConnection(envData)
	if err != nil {
		panic(err)
	}
	err = migration.Migration(db)
	if err != nil {
		panic(errors.New("error running migrations"))
	}
}
