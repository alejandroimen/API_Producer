package repository

import (
	"database/sql"
	"fmt"

	"github.com/alejandroimen/API_Producer/src/users/domain/entities"
)

type UserRepoMySQL struct {
	db *sql.DB
}

func NewCreateUserRepoMySQL(db *sql.DB) *UserRepoMySQL {
	return &UserRepoMySQL{db: db}
}

func (r *UserRepoMySQL) Save(User entities.User) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, User.Name, User.Email, User.Password)
	if err != nil {
		return fmt.Errorf("error insertando User: %w", err)
	}
	return nil
}

func (r *UserRepoMySQL) FindByID(id int) (*entities.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var User entities.User
	if err := row.Scan(&User.ID, &User.Name, &User.Email); err != nil {
		return nil, fmt.Errorf("error buscando el User: %w", err)
	}
	return &User, nil
}

func (r *UserRepoMySQL) FindAll() ([]entities.User, error) {
	query := "SELECT id, name, email, password FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error buscando los Users: %w", err)
	}
	defer rows.Close()

	var Users []entities.User
	for rows.Next() {
		var User entities.User
		if err := rows.Scan(&User.ID, &User.Name, &User.Email, &User.Password); err != nil {
			return nil, err
		}
		Users = append(Users, User)
	}
	return Users, nil
}

func (r *UserRepoMySQL) Update(User entities.User) error {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := r.db.Exec(query, User.Name, User.Email, User.Password, User.ID)
	if err != nil {
		return fmt.Errorf("error actualizando User: %w", err)
	}
	return nil
}

func (r *UserRepoMySQL) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando User: %w", err)
	}
	return nil
}
