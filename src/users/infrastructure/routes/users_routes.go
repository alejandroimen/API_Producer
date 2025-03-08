package routes

import (
	"github.com/alejandroimen/API_Consumer/src/users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupusersRoutes(r *gin.Engine, createusersController *controllers.CreateusersController, getusersController *controllers.GetuserssController, deleteusersController *controllers.DeleteusersController, updateusersController *controllers.UpdateusersController) {
	// Rutas CRUD
	r.POST("/userss", createusersController.Handle)
	r.GET("/userss", getusersController.Handle)
	r.DELETE("/userss/:id", deleteusersController.Handle)
	r.PUT("/userss/:id", updateusersController.Handle)

	// Nuevas rutas para polling en POST
	r.POST("/userss/poll/short", createusersController.ShortPoll)
	r.POST("/userss/poll/long", createusersController.LongPoll)

	// Nuevas rutas para polling en DELETE
	r.DELETE("/userss/poll/short", deleteusersController.ShortPoll)
	r.DELETE("/userss/poll/long", deleteusersController.LongPoll)

	// Nuevas rutas para polling en PUT
	r.PUT("/userss/poll/short", updateusersController.ShortPoll)
	r.PUT("/userss/poll/long", updateusersController.LongPoll)

	// Nuevas rutas para polling en GET
	r.GET("/userss/poll/short", getusersController.ShortPoll)
	r.GET("/userss/poll/long", getusersController.LongPoll)
}
