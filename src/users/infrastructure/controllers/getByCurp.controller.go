package controllers

import (
	"net/http"
	"log"

	"github.com/alejandroimen/API_Producer/src/users/application"
	"github.com/gin-gonic/gin"
)

type GetUserByCURPController struct {
    getUsers      *application.GetUsers
    getUserByCURP *application.GetUserByCURP // Nuevo caso de uso
}

func NewGetUserByCURPController(getUserByCURP *application.GetUserByCURP) *GetUserByCURPController {
    return &GetUserByCURPController{getUserByCURP: getUserByCURP}
}

// Método para obtener un usuario por su CURP
func (gu *GetUserByCURPController) GetUserByCURP(ctx *gin.Context) {
    log.Println("Petición para obtener usuario por CURP, recibida")

    // Obtén el CURP desde los parámetros de la URL
    curp := ctx.Param("curp")

    // Valida que el CURP no sea vacío
    if curp == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "El CURP es obligatorio"})
        return
    }

    // Llama a la lógica de negocio para obtener el usuario
    user, err := gu.getUserByCURP.Run(curp)
    if err != nil {
        log.Printf("Error buscando usuario con CURP %s: %v", curp, err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if user == nil {
        ctx.JSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado"})
        return
    }

    log.Printf("Usuario con CURP %s encontrado", curp)
    ctx.JSON(http.StatusOK, user)
}
