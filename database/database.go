package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DBConn *sql.DB

func ConnectDb() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(0.0.0.0:3306)/books_database"
	client, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
		os.Exit(1)
	}
	log.Println("Connect")
	DBConn = client
}
