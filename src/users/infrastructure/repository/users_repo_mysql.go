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

func (r *UserRepoMySQL) Save(user entities.User) error {
	query := "INSERT INTO users (curp, nombre, apellido, correo) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, user.Curp, user.Nombre, user.Apellido, user.Correo)
	if err != nil {
		return fmt.Errorf("error insertando usuario: %w", err)
	}
	return nil
}

func (r *UserRepoMySQL) FindByCurp(curp string) (*entities.User, error) {
	query := "SELECT curp, nombre, apellido, correo FROM users WHERE curp = ?"
	row := r.db.QueryRow(query, curp)

	var user entities.User
	if err := row.Scan(&user.Curp, &user.Nombre, &user.Apellido, &user.Correo); err != nil {
		return nil, fmt.Errorf("error buscando el usuario: %w", err)
	}
	return &user, nil
}

func (r *UserRepoMySQL) FindAll() ([]entities.User, error) {
	query := "SELECT curp, nombre, apellido, correo FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error buscando los usuarios: %w", err)
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.Curp, &user.Nombre, &user.Apellido, &user.Correo); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepoMySQL) Update(user entities.User) error {
	query := "UPDATE users SET nombre = ?, apellido = ?, correo = ? WHERE curp = ?"
	_, err := r.db.Exec(query, user.Nombre, user.Apellido, user.Correo, user.Curp)
	if err != nil {
		return fmt.Errorf("error actualizando usuario: %w", err)
	}
	return nil
}

func (r *UserRepoMySQL) Delete(curp string) error {
	query := "DELETE FROM users WHERE curp = ?"
	_, err := r.db.Exec(query, curp)
	if err != nil {
		return fmt.Errorf("error eliminando usuario: %w", err)
	}
	return nil
}

func (r *UserRepoMySQL) Login(correo string, password string) (*entities.User, error) {
	query := "SELECT curp, nombre, apellido, correo FROM users WHERE correo = ? AND password = ?"
	row := r.db.QueryRow(query, correo, password)

	var user entities.User
	if err := row.Scan(&user.Curp, &user.Nombre, &user.Apellido, &user.Correo); err != nil {
		return nil, fmt.Errorf("credenciales incorrectas o error en la consulta: %w", err)
	}
	return &user, nil
}
