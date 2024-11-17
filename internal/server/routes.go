package server

import (
	"net/http"

	"github.com/neazk/go_server_skeleton/internal/controllers"
)

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /", controllers.IndexHandler)
	r.HandleFunc("GET /users", controllers.UsersHandler)
	r.HandleFunc("GET /users/{term}", controllers.SearchUsersHandler)
	return r
}
