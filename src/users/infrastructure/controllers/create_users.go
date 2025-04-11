package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/alejandroimen/API_Producer/src/users/application"
	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	CreateUsers *application.CreateUsers
}

func NewCreateUsersController(CreateUsers *application.CreateUsers) *CreateUserController {
	return &CreateUserController{CreateUsers: CreateUsers}
}

func (c *CreateUserController) Handle(ctx *gin.Context) {
	log.Println("Petición de crear un usuario, recibido")

	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando la petición del body: %v", err)
		ctx.JSON(400, gin.H{"error": "petición del body inválida"})
		return
	}
	log.Printf("Creando usuario: Name=%s, email=%s", request.Name, request.Email)

	if err := c.CreateUsers.Run(request.Email, request.Name, request.Password); err != nil {
		log.Printf("Error creando el usuario: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Usuario creado exitosamente")
	ctx.JSON(201, gin.H{"message": "usuario creado exitosamente"})
}

// Controlador para Long Polling
func (c *CreateUserController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}

func waitForNewData() <-chan string {
	newDataChannel := make(chan string)
	go func() {
		time.Sleep(10 * time.Second) // Simula el tiempo hasta que haya nuevos datos
		newDataChannel <- "Datos nuevos disponibles"
	}()
	return newDataChannel
}
