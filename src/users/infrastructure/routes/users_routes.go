package routes

import (
	"github.com/alejandroimen/API_Producer/src/users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, createUserController *controllers.CreateUserController, getUserController *controllers.GetUsersController, deleteUserController *controllers.DeleteUserController, updateUserController *controllers.UpdateUserController) {
	// Rutas CRUD
	r.POST("/users", createUserController.Handle)
	r.GET("/users", getUserController.Handle)
	r.DELETE("/users/:id", deleteUserController.Handle)
	r.PUT("/users/:id", updateUserController.Handle)

	// Nuevas rutas para polling en POST
	r.POST("/users/poll/short", createUserController.ShortPoll)
	r.POST("/users/poll/long", createUserController.LongPoll)

	// Nuevas rutas para polling en DELETE
	r.DELETE("/users/poll/short", deleteUserController.ShortPoll)
	r.DELETE("/users/poll/long", deleteUserController.LongPoll)

	// Nuevas rutas para polling en PUT
	r.PUT("/users/poll/short", updateUserController.ShortPoll)
	r.PUT("/users/poll/long", updateUserController.LongPoll)

	// Nuevas rutas para polling en GET
	r.GET("/users/poll/short", getUserController.ShortPoll)
	r.GET("/users/poll/long", getUserController.LongPoll)
}
