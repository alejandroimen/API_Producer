package main

import (
	"log"

	"github.com/alejandroimen/API_Producer/src/core"
	usersApp "github.com/alejandroimen/API_Producer/src/users/application"
	usersController "github.com/alejandroimen/API_Producer/src/users/infrastructure/controllers"
	usersRepo "github.com/alejandroimen/API_Producer/src/users/infrastructure/repository"
	usersRoutes "github.com/alejandroimen/API_Producer/src/users/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conexión a MySQL
	db, err := core.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Repositorios
	usersRepository := usersRepo.NewCreateUserRepoMySQL(db)

	// Casos de uso para users
	createUsers := usersApp.NewCreateUsers(usersRepository)
	getUsers := usersApp.NewGetUsers(usersRepository)
	deleteUsers := usersApp.NewDeleteUsers(usersRepository)
	updateUsers := usersApp.NewUpdateUsers(usersRepository)

	// Controladores para users
	createusersController := usersController.NewCreateUsersController(createUsers)
	getusersController := usersController.NewUsersController(getUsers)
	deleteusersController := usersController.NewDeleteUsersController(deleteUsers)
	updateusersController := usersController.NewUpdateUsersController(updateUsers)

	// Configuración del enrutador de Gin
	r := gin.Default()

	// Configurar rutas de users
	usersRoutes.SetupUserRoutes(r, createusersController, getusersController, deleteusersController, updateusersController)

	// Iniciar servidor
	log.Println("server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
