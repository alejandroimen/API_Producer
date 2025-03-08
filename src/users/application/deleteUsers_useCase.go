package application

import (
	"fmt"

	"github.com/alejandroimen/API_Consumer/src/users/domain/repository"
)

type DeleteUsers struct {
	repo repository.usersRepository
}

func NewDeleteUsers(repo repository.usersRepository) *DeleteUsers {
	return &DeleteUsers{repo: repo}
}

func (du *DeleteUsers) Run(id int) error {
	_, err := du.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("users no encontrado: %w", err)
	}

	if err := du.repo.Delete(id); err != nil {
		return fmt.Errorf("error eliminando el usuairo: %w", err)
	}
	return nil
}
