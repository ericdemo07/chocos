package server

import (
	"example/hello/app"
	"example/hello/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func routes(dependency *app.Dependency) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", handlers.ProductsHandler(dependency.ProductService)).
		Methods(http.MethodGet).
		Name("GetProducts").
		Schemes("http")

	return router
}
