package db

import (
	"database/sql"
	"fmt"
	"log"
	// _ "github.com/go-sql-driver/mysql"
)

func testConnection(db *sql.DB) {
	// Test the connection to the database
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func connect() (*sql.DB, error) {
	// MySQL connection configuration
	// "mysql", "maath:_14326@tcp(127.0.0.1:3306)/db_go_react"
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		return nil, err
	}

	testConnection(db)

	return db, nil
}

func executeQuery(db *sql.DB, query string) (*sql.Rows, error) {
	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func sample() {
	// Connect to MySQL
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	// Execute a sample query
	rows, err := executeQuery(db, "SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	processingResultsSample(rows)

	// Close the database connection
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func processingResultsSample(rows *sql.Rows) {
	// Process the query results
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID:", id, "Name:", name)
	}
}
