package handlers

import (
	"example/hello/products"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ProductsHandler(service products.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		ctx := r.Context()
		service.GetData(ctx)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Category: %v\n", vars["category"])
	}
}
