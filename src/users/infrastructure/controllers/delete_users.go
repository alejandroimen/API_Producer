package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/alejandroimen/API_Consumer/src/users/application"
	"github.com/gin-gonic/gin"
)

type DeleteusersController struct {
	deleteusers *application.Deleteusers
}

func NewDeleteusersController(deleteusers *application.Deleteusers) *DeleteusersController {
	return &DeleteusersController{deleteusers: deleteusers}
}

func (du *DeleteusersController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID de user inválido"})
		return
	}

	if err := du.deleteusers.Run(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "user eliminado correctamente"})
}

// Controlador para Short Polling
func (du *DeleteusersController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (du *DeleteusersController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
