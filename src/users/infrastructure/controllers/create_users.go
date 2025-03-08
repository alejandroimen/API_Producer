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
	log.Println("Petición de crear un usuario recibida")

	var request struct {
		Curp      string `json:"curp"`
		Nombre    string `json:"nombre"`
		Apellido  string `json:"apellido"`
		Correo    string `json:"correo"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando la petición del body: %v", err)
		ctx.JSON(400, gin.H{"error": "petición del body inválida"})
		return
	}
	log.Printf("Creando usuario: curp=%s, nombre=%s, apellido=%s, correo=%s", request.Curp, request.Nombre, request.Apellido, request.Correo)

	if err := c.CreateUsers.Run(request.Curp, request.Nombre, request.Apellido, request.Correo); err != nil {
		log.Printf("Error creando el usuario: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Println("Usuario creado exitosamente")
	ctx.JSON(201, gin.H{"message": "usuario creado exitosamente"})
}

// Controlador para Short Polling
func (c *CreateUserController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
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
