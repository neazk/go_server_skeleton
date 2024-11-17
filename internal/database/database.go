package database

import (
	"database/sql"
	"log"
	"os"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

var dbConn *sql.DB

// var (
// 	dbname   = os.Getenv("DB_DATABASE")
// 	password = os.Getenv("DB_PASSWORD")
// 	username = os.Getenv("DB_USERNAME")
// 	port     = os.Getenv("DB_PORT")
// 	host     = os.Getenv("DB_HOST")
// )

func DbConnection() *sql.DB {
	if dbConn == nil {
		db, err := sql.Open("sqlite3", os.Getenv("DB_URL"))
		if err != nil {
			log.Fatal("error opening database: ", err)
		}

		// db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))
		// if err != nil {
		// 	log.Fatal("error opening database: ", err)
		// }
		// db.SetConnMaxLifetime(time.Minute * 3)
		// db.SetMaxOpenConns(10)
		// db.SetMaxIdleConns(10)
		dbConn = db
	}
	return dbConn
}
