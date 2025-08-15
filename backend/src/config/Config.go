package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host,Port,DBName,Username,Password string
	RHost,RPort,RUsername,RPassword string
	JwtKey string
}

func NewConfig() *Env {
	err := godotenv.Load("./.env");
	if err != nil {
		log.Fatalln("Error Load .env file")
	}

	return &Env{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),

		// redis
		RHost: os.Getenv("RDB__HOST"),
		RPort: os.Getenv("RDB_PORT"),
		RUsername: os.Getenv("RDB_USERNAME"),
		RPassword: os.Getenv("RDB_PASSWORD"),

		// jwt
		JwtKey: os.Getenv("JWT_SECRET_KEY"),
	}
}