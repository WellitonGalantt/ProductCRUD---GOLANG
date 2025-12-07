package main

import (
	"productcrud/controller"
	"productcrud/db"
	"productcrud/middleware"
	"productcrud/repository"
	"productcrud/usecase"

	"github.com/gin-gonic/gin"

	"log"

	"github.com/joho/godotenv"
)

//go mod tidy

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Aviso: .env não encontrado, usando variáveis do sistema")
	}

	server := gin.Default()

	dbConection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConection)
	UserRepository := repository.NewUserRepository(dbConection)

	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	UserUsecase := usecase.NewUserUsecase(UserRepository)

	ProductController := controller.NewProductController(ProductUsecase)
	UserController := controller.NewUserController(UserUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"messsage": "pong",
		})
	})

	//proteger as rotas:
	// Todas que tiverem /api antes devem passar o header/token
	api := server.Group("/api")
	api.Use(middleware.AuthMiddleware())

	api.GET("/products", ProductController.GetProducts)
	api.POST("/product", ProductController.CreateProduct)
	api.GET("/product/:pdId", ProductController.GetProductById)
	api.PUT("/product/:pdId", ProductController.UpdateProduct)

	server.POST("/user/register", UserController.RegisterUser)
	server.POST("/user/login", UserController.LoginUser)

	server.Run(":8000")
}
