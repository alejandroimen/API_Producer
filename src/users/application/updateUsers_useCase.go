package application

import (
	"fmt"

	_ "github.com/alejandroimen/API_Producer/src/users/domain/entities"
	"github.com/alejandroimen/API_Producer/src/users/domain/repository"
)

type UpdateUser struct {
	repo repository.UserRepository
}

func NewUpdateUser(repo repository.UserRepository) *UpdateUser {
	return &UpdateUser{repo: repo}
}

func (us *UpdateUser) Run(id int, name string, email string) error {
	user, err := us.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("usuario no encontrado: %w", err)
	}

	//actualizo los campos del usuario:
	user.Name = name
	user.Email = email

	//guardo los cambios en el repositorio:
	if err := us.repo.Update(*user); err != nil {
		return fmt.Errorf("error actualizando el usuario: %w", err)
	}

	return nil
}
