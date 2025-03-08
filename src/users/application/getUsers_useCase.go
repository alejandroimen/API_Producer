package application

import (
	"github.com/alejandroimen/API_Consumer/src/users/domain/repository"
	"github.com/alejandroimen/API_Consumer/src/users/domain/entities"
)

type GetUsers struct {
	repo repository.usersRepository
}

func NewGetUsers(repo repository.usersRepository) *GetUsers {
	return &GetUsers{repo: repo}
}

func (gu *GetUsers) Run() ([]entities.users, error) {
	userss, err := gu.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return userss, nil
}
