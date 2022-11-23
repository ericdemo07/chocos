package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
