package application

import (
	"fmt"

	"github.com/alejandroimen/API_Producer/src/users/domain/repository"
)

type DeleteUser struct {
	repo repository.UserRepository
}

func NewDeleteUsers(repo repository.UserRepository) *DeleteUser {
	return &DeleteUser{repo: repo}
}

func (du *DeleteUser) Run(curp string) error {
	_, err := du.repo.FindByCurp(curp)
	if err != nil {
		return fmt.Errorf("user no encontrado: %w", err)
	}

	if err := du.repo.Delete(curp); err != nil {
		return fmt.Errorf("error eliminando el usuairo: %w", err)
	}
	return nil
}
