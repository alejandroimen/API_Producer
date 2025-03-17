package application

import (
	"fmt"

	_ "github.com/alejandroimen/API_Producer/src/users/domain/entities"
	"github.com/alejandroimen/API_Producer/src/users/domain/repository"
)

type UpdateUser struct {
	repo repository.UserRepository
}

func NewUpdateUsers(repo repository.UserRepository) *UpdateUser {
	return &UpdateUser{repo: repo}
}

func (us *UpdateUser) Run(curp string, nombre string, correo string, apellido string) error {
	user, err := us.repo.FindByCurp(curp)
	if err != nil {
		return fmt.Errorf("usuario no encontrado: %w", err)
	}

	//actualizo los campos del usuario:
	user.Curp = curp
	user.Nombre = nombre
	user.Apellido = apellido
	user.Correo = correo

	//guardo los cambios en el repositorio:
	if err := us.repo.Update(*user); err != nil {
		return fmt.Errorf("error actualizando el usuario: %w", err)
	}

	return nil
}
