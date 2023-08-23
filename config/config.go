package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDatabase connects to a MySQL database and returns a *sql.DB object
func ConnectDatabase(host, dbName, user, password string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	// Set other connection settings if needed
	db.SetMaxOpenConns(10) // Example: Limit the number of open connections

	// Ping the database to check connectivity
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
