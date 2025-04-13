package application

import (
	"fmt"
	"log"

	"github.com/alejandroimen/API_Producer/src/users/domain/entities"
	"github.com/alejandroimen/API_Producer/src/users/domain/repository"
	"github.com/alejandroimen/API_Producer/src/users/domain/services"
)

// Contiene un campo de repo de tipo repository.user... siendo esto una inyección de dependencias
type CreateUsers struct {
	repo repository.UserRepository
	rab services.RabbitMQService
}

// constructor de createusers, que recibe un repositorio como parametro y lo asigna al campo repo. siendo configurable
func NewCreateUser(repo repository.UserRepository, rab services.RabbitMQService) *CreateUsers {
	return &CreateUsers{repo: repo, rab: rab}
}

func (cu *CreateUsers) Run(curp string, name string, lastname string, phone string, email string) error {
	user := entities.User{CURP: curp, Name: name, Lastname: lastname, Phone: phone, Email: email}
	id, err := cu.repo.Save(user); 
	if err != nil {
		return fmt.Errorf("error al guardar el usuario: %w", err)
	}

	err = cu.rab.PublishCreatedUser(id)
	if err != nil {
		return fmt.Errorf("error al publicar evento en la cola: %w", err)
	}

	log.Printf("✅ Evento 'cita.created' publicado para el usuario %s", user.Name)
	return nil
}
