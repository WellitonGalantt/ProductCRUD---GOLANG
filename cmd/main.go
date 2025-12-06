package main

import (
	"productcrud/controller"
	"productcrud/db"
	"productcrud/repository"
	"productcrud/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
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

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:pdId", ProductController.GetProductById)
	server.PUT("/product/:pdId", ProductController.UpdateProduct)

	server.POST("/user/register", UserController.RegisterUser)
	server.POST("/user/login", UserController.LoginUser)

	server.Run(":8000")
}
