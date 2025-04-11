package infraestructure

import (
	"database/sql"
	"github.com/gin-gonic/gin"
    "os"
	userApp "github.com/alejandroimen/API_Producer/src/users/application"
    userController "github.com/alejandroimen/API_Producer/src/users/infraestructure/controllers"
    userRepo "github.com/alejandroimen/API_Producer/src/users/infraestructure/repository"
    userRoutes "github.com/alejandroimen/API_Producer/src/users/infraestructure/routes"
    "github.com/alejandroimen/API_Producer/src/users/infraestructure/services"
)

func InitUsersDependencies(Engine *gin.Engine, db *sql.DB){
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET no está configurado en las variables de entorno")
	}
    
	userRepository := userRepo.NewCreateUserRepoMySQL(db)

	createUser := userApp.NewCreateUser(userRepository)
    getUsers := userApp.NewGetUsers(userRepository)
    deleteUsers := userApp.NewDeleteUser(userRepository)
    updateUsers := userApp.NewUpdateUser(userRepository)
    loginUser := userApp.NewLoginUser(userRepository)     

	createUserController := userController.NewCreateUserController(createUser)
    getUserController := userController.NewUsersController(getUsers)
    deleteUserController := userController.NewDeleteUserController(deleteUsers)
    updateUserController := userController.NewUpdateUserController(updateUsers)
    loginUserController := userController.NewLoginUserController(loginUser)

	userRoutes.SetupUserRoutes(Engine, createUserController, loginUserController, getUserController, deleteUserController, updateUserController)

}