package configuration

import "os"

type EnvData struct {
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	Port      string
	JWTSecret string
	JWTTime   string
}

func NewEnvData() EnvData {
	return EnvData{
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBName:    os.Getenv("DB_NAME"),
		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTTime:   os.Getenv("JWT_TIME"),
	}
}
