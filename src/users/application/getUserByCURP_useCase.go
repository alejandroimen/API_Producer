package application

import (
	"github.com/alejandroimen/API_Producer/src/users/domain/repository"
	"github.com/alejandroimen/API_Producer/src/users/domain/entities"

)


type GetUserByCURP struct {
    repo repository.UserRepository // Interfaz del repositorio
}

func NewGetUserByCURP(repo repository.UserRepository) *GetUserByCURP {
    return &GetUserByCURP{repo: repo}
}

func (uc *GetUserByCURP) Run(curp string) (*entities.User, error) {
    return uc.repo.FindByCurp(curp)
}