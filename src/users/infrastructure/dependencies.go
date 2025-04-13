package infrastructure

import (
	"database/sql"
    "log"

	userApp "github.com/alejandroimen/API_Producer/src/users/application"
	userController "github.com/alejandroimen/API_Producer/src/users/infrastructure/controllers"
	userRepo "github.com/alejandroimen/API_Producer/src/users/infrastructure/repository"
	userRoutes "github.com/alejandroimen/API_Producer/src/users/infrastructure/routes"
    "github.com/alejandroimen/API_Producer/src/users/infrastructure/adapters"
	"github.com/gin-gonic/gin"
)

func InitUsersDependencies(Engine *gin.Engine, db *sql.DB) {


    // Inicializar RabbitMQ
    rabbitMQService, err := adapters.NewRabbitMQAdapter("amqp://rabbit:rabbit@35.170.173.77:5672/vh")
    if err != nil {
        log.Fatalf("Error inicializando RabbitMQ: %s", err)
    }
    rabbitMQService.StartConsumingCitas()

    defer rabbitMQService.Close() // Garantizamos el cierre al terminar

	userRepository := userRepo.NewCreateUserRepoMySQL(db)

	createUser := userApp.NewCreateUser(userRepository, rabbitMQService)
	getUsers := userApp.NewGetUsers(userRepository)
	deleteUsers := userApp.NewDeleteUser(userRepository)
	updateUsers := userApp.NewUpdateUser(userRepository)
    getUserByCURP := userApp.NewGetUserByCURP(userRepository)

	createUserController := userController.NewCreateUsersController(createUser)
	getUserController := userController.NewUsersController(getUsers)
	deleteUserController := userController.NewDeleteUsersController(deleteUsers)
	updateUserController := userController.NewUpdateUsersController(updateUsers)
    getUserByCURPController := userController.NewGetUserByCURPController(getUserByCURP)

	userRoutes.SetupUserRoutes(Engine, createUserController, getUserController, getUserByCURPController, deleteUserController, updateUserController)

}
