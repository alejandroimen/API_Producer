package repository

import "github.com/alejandroimen/API_Producer/src/users/domain/entities"

type UserRepository interface {
	Save(user entities.User) error
	FindByCurp(id string) (*entities.User , error)
	FindAll() ([]entities.User, error)
	Update(user entities.User) error
	Delete(curp string) error
}
