package core

// Aquí simple y sencillamente la configuración con MySQL

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection() (*sql.DB, error) {
	dsn := "root:helloAtu21.@tcp(127.0.0.1:3306)/arquitectura"
	//dsn := "root:chocolate200201614@tcp(127.0.0.1:3306)/hexagonality"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MySQL: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging MySQL: %w", err)
	}
	return db, nil
}
