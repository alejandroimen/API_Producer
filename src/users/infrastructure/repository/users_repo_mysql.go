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

func (r *UserRepoMySQL) Save(User entities.User) (int, error) {
	query := "INSERT INTO users (curp, name, lastname, phone, email) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, User.CURP, User.Name, User.Lastname, User.Phone, User.Email)
	if err != nil {
		return 0, fmt.Errorf("error insertando User: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error al obtener el ID del usuario: %w", err)
	}

	return int(id), nil
}


func (r *UserRepoMySQL) FindByID(id int) (*entities.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var User entities.User
	if err := row.Scan(&User.ID, User.CURP, &User.Name, &User.Lastname, &User.Phone, &User.Email); err != nil {
		return nil, fmt.Errorf("error buscando el User: %w", err)
	}
	return &User, nil
}

func (r *UserRepoMySQL) FindByCurp(curp string) (*entities.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.db.QueryRow(query, curp)

	var User entities.User
	if err := row.Scan(&User.ID, User.CURP, &User.Name, &User.Lastname, &User.Phone, &User.Email); err != nil {
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
		if err := rows.Scan(&User.ID, &User.CURP, &User.Name, &User.Lastname, &User.Phone, &User.Email); err != nil {
			return nil, err
		}
		Users = append(Users, User)
	}
	return Users, nil
}

func (r *UserRepoMySQL) Update(User entities.User) error {
	query := "UPDATE users SET curp = ?, name = ?, lastname = ?, phone = ? email = ? WHERE id = ?"
	_, err := r.db.Exec(query, User.CURP, User.Name, User.Lastname, User.Phone, User.Email, User.ID)
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
