package configuration

import (
	"os"
	"strconv"
)

type EnvData struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPass       string
	DBName       string
	Port         int
	JWTSecret    string
	JWTTime      int
	InProduction bool
}

func NewEnvData() EnvData {
	jwtTime, _ := strconv.Atoi(os.Getenv("JWT_TIME"))
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return EnvData{
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPass:       os.Getenv("DB_PASS"),
		DBName:       os.Getenv("DB_NAME"),
		Port:         port,
		JWTSecret:    os.Getenv("JWT_SECRET"),
		JWTTime:      jwtTime,
		InProduction: os.Getenv("IN_PRODUCTION") == "true",
	}
}
