package application

import (
	"fmt"

	"github.com/alejandroimen/API_Consumer/src/users/domain/repository"
)

type UpdateUsers struct {
	repo repository.usersRepository
}

func NewUpdateUsers(repo repository.usersRepository) *UpdateUsers {
	return &UpdateUsers{repo: repo}
}

func (us *UpdateUsers) Run(id int, name string, email string, password string) error {
	users, err := us.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("user no encontrado: %w", err)
	}

	//actualizo los campos del user:
	users.Name = name
	users.Email = email
	users.Password = password

	//guardo los cambios en el repositorio:
	if err := us.repo.Update(*users); err != nil {
		return fmt.Errorf("error actualizando el user: %w", err)
	}

	return nil
}
