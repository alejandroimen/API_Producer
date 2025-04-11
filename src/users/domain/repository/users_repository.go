package repository

import "github.com/alejandroimen/API_Producer/src/users/domain/entities"

type UserRepository interface {
	Save(user entities.User) error
	FindByID(id int) (*entities.User, error)
	FindAll() ([]entities.User, error)
	Update(user entities.User) error
	Delete(id int) error
	FindByCurp(email string) (*entities.User, error)
}