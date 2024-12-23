package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	err error
)

func Init() *sql.DB {
    // Define MySQL connection details

	// MySQL connection string
	dsn := "root:abc123@tcp(go_basic_db:3306)/go_basic_db?charset=utf8mb4"

	// Connect to MySQL
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to connect to MySQL: %v", err)
	}
	// defer db.Close()

	fmt.Println("database: ==================", db)


	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	fmt.Println("Successfully connected to the database!")

	return db
}

func GetDB() *sql.DB {
	return db
}
