package repository

import (
	"database/sql"
	"fmt"

	"github.com/alejandroimen/API_Consumer/src/users/domain/entities"
)

type usersRepoMySQL struct {
	db *sql.DB
}

func NewCreateusersRepoMySQL(db *sql.DB) *usersRepoMySQL {
	return &usersRepoMySQL{db: db}
}

func (r *usersRepoMySQL) Save(users entities.users) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, users.Name, users.Email, users.Password)
	if err != nil {
		return fmt.Errorf("error insertando users: %w", err)
	}
	return nil
}

func (r *usersRepoMySQL) FindByID(id int) (*entities.users, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var users entities.users
	if err := row.Scan(&users.ID, &users.Name, &users.Email); err != nil {
		return nil, fmt.Errorf("error buscando el users: %w", err)
	}
	return &users, nil
}

func (r *usersRepoMySQL) FindAll() ([]entities.users, error) {
	query := "SELECT id, name, email, password FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error buscando los users: %w", err)
	}
	defer rows.Close()

	var users []entities.users
	for rows.Next() {
		var users entities.users
		if err := rows.Scan(&users.ID, &users.Name, &users.Email, &users.Password); err != nil {
			return nil, err
		}
		users = append(users, users)
	}
	return users, nil
}

func (r *usersRepoMySQL) Update(users entities.users) error {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := r.db.Exec(query, users.Name, users.Email, users.Password, users.ID)
	if err != nil {
		return fmt.Errorf("error actualizando users: %w", err)
	}
	return nil
}

func (r *usersRepoMySQL) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando users: %w", err)
	}
	return nil
}
