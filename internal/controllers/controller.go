package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/neazk/go_server_skeleton/internal/database"
)

var db *sql.DB = database.DbConnection()

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Food index")
}

func SearchUsersHandler(w http.ResponseWriter, r *http.Request) {
	term := r.PathValue("term")
	res := db.QueryRow(`SELECT name FROM users WHERE id = ?`, term)
	fmt.Fprint(w, res)
}
