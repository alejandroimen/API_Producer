package controllers

import (
	"net/http"
	"time"

	"github.com/alejandroimen/API_Producer/src/users/application"
	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	deleteUser *application.DeleteUser
}

func NewDeleteUsersController(deleteUser *application.DeleteUser) *DeleteUserController {
	return &DeleteUserController{deleteUser: deleteUser}
}

func (du *DeleteUserController) Handle(ctx *gin.Context) {
	id := ctx.Param("curp")

	if err := du.deleteUser.Run(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "usuario eliminado correctamente"})
}

// Controlador para Short Polling
func (du *DeleteUserController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (du *DeleteUserController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
