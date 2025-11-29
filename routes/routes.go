package routes

import (
	"net/http"
	"testMKP/controllers"
)

func SetupRoutes(userController *controllers.UserController, productController *controllers.ProductController) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/create", userController.Create)
	mux.HandleFunc("/users/list", userController.FindAll)
	mux.HandleFunc("/users/detail", userController.FindById)
	mux.HandleFunc("/users/update", userController.Update)
	mux.HandleFunc("/users/delete", userController.Delete)

	mux.HandleFunc("/products/create", productController.Create)
	mux.HandleFunc("/products/list", productController.FindAll)
	mux.HandleFunc("/products/detail", productController.FindById)
	mux.HandleFunc("/products/update", productController.Update)
	mux.HandleFunc("/products/delete", productController.Delete)

	return mux
}
