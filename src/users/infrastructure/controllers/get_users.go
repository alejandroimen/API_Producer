package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/alejandroimen/API_Producer/src/users/application"
	"github.com/gin-gonic/gin"
)

type GetUsersController struct {
	getUsers *application.GetUsers
}

func NewUsersController(getUsers *application.GetUsers) *GetUsersController {
	return &GetUsersController{getUsers: getUsers}
}

func (gu *GetUsersController) Handle(ctx *gin.Context) {
	log.Println("Petición de listar todos los usuarios, recibido")

	user, err := gu.getUsers.Run()
	if err != nil {
		log.Printf("Error buscando usuarios")
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Retornando %d usuarios", len(user))
	ctx.JSON(200, user)
}
func (c *GetUsersController) ShortPoll(ctx *gin.Context) {
	// Obtener los productos (esto simula si hay cambios o no)
	products, err := c.getUsers.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		// No hay productos (o cambios)
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
		return
	}

	// Devolver productos (o cambios detectados)
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Datos actualizados",
		"products": products,
	})
}

// Controlador para Long Polling
func (gu *GetUsersController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
