package application

import (
	"github.com/alejandroimen/API_Producer/src/users/domain/entities"
	"github.com/alejandroimen/API_Producer/src/users/domain/repository"
)

type GetUsers struct {
	repo repository.UserRepository
}

func NewGetUsers(repo repository.UserRepository) *GetUsers {
	return &GetUsers{repo: repo}
}

func (gu *GetUsers) Run() ([]entities.User, error) {
	users, err := gu.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
