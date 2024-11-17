package server

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func NewServer() *http.Server {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := LoggerMiddleware(AuthMiddleware(NewRouter()))
	return &http.Server{Addr: ":" + os.Getenv("PORT"), Handler: r}
}
