package main

import (
	"log"

	"github.com/alejandroimen/API_Consumer/src/core"
	productApp "github.com/alejandroimen/API_Consumer/src/products/application"
	productController "github.com/alejandroimen/API_Consumer/src/products/infrastructure/controllers"
	productRepo "github.com/alejandroimen/API_Consumer/src/products/infrastructure/repository"
	productRoutes "github.com/alejandroimen/API_Consumer/src/products/infrastructure/routes"
	usersApp "github.com/alejandroimen/API_Consumer/src/users/application"
	usersController "github.com/alejandroimen/API_Consumer/src/users/infrastructure/controllers"
	usersRepo "github.com/alejandroimen/API_Consumer/src/users/infrastructure/repository"
	usersRoutes "github.com/alejandroimen/API_Consumer/src/users/infrastructure/routes"
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
	productRepository := productRepo.NewProductRepoMySQL(db)
	usersRepository := usersRepo.NewCreateusersRepoMySQL(db)

	// Casos de uso para productos
	createProduct := productApp.NewCreateProduct(productRepository)
	getProducts := productApp.NewGetProducts(productRepository)
	updateProduct := productApp.NewUpdateProduct(productRepository)
	deleteProduct := productApp.NewDeleteProduct(productRepository)

	// Casos de uso para users
	createusers := usersApp.NewCreateusers(usersRepository)
	getuserss := usersApp.NewGetuserss(usersRepository)
	deleteuserss := usersApp.NewDeleteusers(usersRepository)
	updateuserss := usersApp.NewUpdateusers(usersRepository)

	// Controladores para productos
	createProductController := productController.NewCreateProductController(createProduct)
	getProductsController := productController.NewGetProductsController(getProducts)
	updateProductController := productController.NewUpdateProductController(updateProduct)
	deleteProductController := productController.NewDeleteProductController(deleteProduct)

	// Controladores para users
	createusersController := usersController.NewCreateusersController(createusers)
	getusersController := usersController.NewuserssController(getuserss)
	deleteusersController := usersController.NewDeleteusersController(deleteuserss)
	updateusersController := usersController.NewUpdateusersController(updateuserss)

	// Configuración del enrutador de Gin
	r := gin.Default()

	// Configurar rutas de productos
	productRoutes.SetupProductRoutes(r, createProductController, getProductsController, updateProductController, deleteProductController)

	// Configurar rutas de users
	usersRoutes.SetupusersRoutes(r, createusersController, getusersController, deleteusersController, updateusersController)

	// Iniciar servidor
	log.Println("server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
