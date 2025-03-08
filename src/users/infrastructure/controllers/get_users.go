package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/alejandroimen/API_Consumer/src/users/application"
	"github.com/gin-gonic/gin"
)

type GetuserssController struct {
	getuserss *application.Getuserss
}

func NewuserssController(getuserss *application.Getuserss) *GetuserssController {
	return &GetuserssController{getuserss: getuserss}
}

func (gu *GetuserssController) Handle(ctx *gin.Context) {
	log.Println("Petición de listar todos los users, recibido")

	users, err := gu.getuserss.Run()
	if err != nil {
		log.Printf("Error buscando users")
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Retornando %d users", len(users))
	ctx.JSON(200, users)
}
func (c *GetuserssController) ShortPoll(ctx *gin.Context) {
	// Obtener los productos (esto simula si hay cambios o no)
	products, err := c.getuserss.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalserverError, gin.H{"error": err.Error()})
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
func (gu *GetuserssController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
