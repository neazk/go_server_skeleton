package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neazk/go_server_skeleton/internal/database"
)

var db *sql.DB = database.DbConnection()

func IndexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func UsersHandler(c *gin.Context) {
	c.String(http.StatusOK, "Users Page")
}

func SearchUsersHandler(c *gin.Context) {
	term := c.Param("term")
	res := db.QueryRow(`SELECT name FROM users WHERE id = ?`, term)
	c.JSON(http.StatusOK, res)
}
