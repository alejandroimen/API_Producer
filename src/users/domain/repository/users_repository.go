package repository

import "github.com/alejandroimen/API_Consumer/src/users/domain/entities"

type usersRepository interface {
	Save(users entities.users) error
	FindByID(id int) (*entities.users, error)
	FindAll() ([]entities.users, error)
	Update(users entities.users) error
	Delete(id int) error
}
