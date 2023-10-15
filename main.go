package main

import (
	"go-api/helpers"
	"go-api/server"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	helpers.InitJWTKey([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func main() {
	srv := server.NewRestServer()
	srv.Run()
}
