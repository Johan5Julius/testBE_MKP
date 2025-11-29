package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testMKP/belajar/helper"
	"testMKP/controllers"
	"testMKP/database"
	"testMKP/repositories"
	"testMKP/routes"
	"testMKP/services"
)

func main() {
	db, err := database.GetConnection()
	helper.PanicIfError(err)
	defer db.Close()

	mongoClient, err := database.GetConnectionMongoDB()
	helper.PanicIfError(err)
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
		fmt.Println("MongoDB connection closed")
	}()
	
	mongoDB := database.GetDatabase(mongoClient, "productdb")

	userRepository := repositories.NewUserRepositoryImpl(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	productRepository := repositories.NewProductRepository(mongoDB)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	router := routes.SetupRoutes(userController, productController)

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
