package routes

import (
	"github.com/alejandroimen/API_Producer/src/users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, 
						createUserController *controllers.CreateUserController, 
						getUserController *controllers.GetUsersController, 
						getByCurp *controllers.GetUserByCURPController,
						deleteUserController *controllers.DeleteUserController, 
						updateUserController *controllers.UpdateUserController) {
	// Rutas CRUD
	r.POST("/users", createUserController.Handle)
	r.GET("/users", getUserController.Handle)
	r.DELETE("/users/:id", deleteUserController.Handle)
	r.PUT("/users/:id", updateUserController.Handle)

	
}
