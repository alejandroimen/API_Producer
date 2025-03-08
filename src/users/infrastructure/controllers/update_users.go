package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/alejandroimen/API_Consumer/src/users/application"
	"github.com/gin-gonic/gin"
)

type UpdateusersController struct {
	updateusers *application.Updateusers
}

func NewUpdateusersController(updateusers *application.Updateusers) *UpdateusersController {
	return &UpdateusersController{updateusers: updateusers}
}

func (update *UpdateusersController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID de user inválido"})
		return
	}

	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "petición del body inválida"})
		return
	}

	if err := update.updateusers.Run(id, request.Email, request.Name, request.Password); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "user actualizado correctamente"})
}

// Controlador para Short Polling
func (update *UpdateusersController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (update *UpdateusersController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
