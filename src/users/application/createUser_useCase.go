package application

import (
	"fmt"

	"github.com/alejandroimen/API_Consumer/src/users/domain/repository"
	"github.com/alejandroimen/API_Consumer/src/users/domain/entities"
)

// Contiene un campo de repo de tipo repository.user... siendo esto una inyección de dependencias
type CreateUsers struct {
	repo repository.usersRepository
}

// constructor de createusers, que recibe un repositorio como parametro y lo asigna al campo repo. siendo configurable
func NewCreateUser(repo repository.usersRepository) *CreateUsers {
	return &CreateUsers{repo: repo}
}

func (cu *CreateUsers) Run(name string, email string, password string) error {
	user := entities.User{Name: name, Email: email, Password: password}
	if err := cu.repo.Save(user); err != nil {
		return fmt.Errorf("error al guardar el user: %w", err)
	}
	return nil
}
