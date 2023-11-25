package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func loadConfig() (Config, error) {
	// Load the configuration from a file or any other source
	// Implement your logic here to load the config
	config := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "your_password",
		DBName:   "your_database",
	}

	return config, nil
}

func connectDB() (*sql.DB, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	testConnection(db)

	return db, nil
}

func executePsqlQuery(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func Sample() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT * FROM users"
	err = executePsqlQuery(db, query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Query executed successfully!")
}
