package server

import (
	"example/hello/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", handlers.ProductsHandler).
		Methods(http.MethodGet).
		Name("GetProducts").
		Schemes("http")

	return router
}
