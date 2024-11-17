package main

import (
	"log"

	"github.com/neazk/go_server_skeleton/internal/server"
)

func main() {
	s := server.NewServer()
	log.Println("Listening on port", s.Addr)
	s.ListenAndServe()
}
