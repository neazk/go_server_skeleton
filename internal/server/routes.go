package server

import (
	"github.com/gin-gonic/gin"
	"github.com/neazk/go_server_skeleton/internal/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.IndexHandler)
	r.GET("/users", controllers.UsersHandler)
	r.GET("/search/:term", controllers.SearchUsersHandler)
	return r
}
