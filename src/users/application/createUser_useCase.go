package application

import (
	"fmt"

	"github.com/alejandroimen/API_Producer/src/users/domain/entities"
	"github.com/alejandroimen/API_Producer/src/users/domain/repository"
)

// Contiene un campo de repo de tipo repository.user... siendo esto una inyección de dependencias
type CreateUsers struct {
	repo repository.UserRepository
}

// constructor de createusers, que recibe un repositorio como parametro y lo asigna al campo repo. siendo configurable
func NewCreateUsers(repo repository.UserRepository) *CreateUsers {
	return &CreateUsers{repo: repo}
}

func (cu *CreateUsers) Run(curp string, nombre string, apellido string, correo string) error {
	user := entities.User{Curp: curp, Nombre: nombre, Apellido: apellido, Correo: correo}
	if err := cu.repo.Save(user); err != nil {
		return fmt.Errorf("error al guardar el usuario: %w", err)
	}
	return nil
}
