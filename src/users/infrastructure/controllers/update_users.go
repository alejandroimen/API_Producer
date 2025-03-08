package controllers

import (
	"net/http"
	"time"

	"github.com/alejandroimen/API_Producer/src/users/application"
	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUser *application.UpdateUser
}

func NewUpdateUsersController(updateUser *application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{updateUser: updateUser}
}

func (update *UpdateUserController) Handle(ctx *gin.Context) {
	curp := ctx.Param("curp")

	var request struct {
		Nombre     string `json:"name"`
		Apellido    string `json:"email"`
		Correo string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "petición del body inválida"})
		return
	}

	if err := update.updateUser.Run(curp, request.Nombre, request.Apellido, request.Correo); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "usuario actualizado correctamente"})
}

// Controlador para Short Polling
func (update *UpdateUserController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (update *UpdateUserController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
