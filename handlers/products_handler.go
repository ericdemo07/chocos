package handlers

import (
	"encoding/json"
	"example/hello/products"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ProductsHandler(service products.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		//ctx := r.Context()
		//service.GetData(ctx)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Println(w, "Category: %v\n", vars["category"])

		responseAsByteArray, _ := json.Marshal(productList)

		w.Write(responseAsByteArray)
	}
}

type brand struct {
	ID               string   `json:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"shortDescription"`
	Description      string   `json:"description"`
	Images           []string `json:"images"`
}

// this will be a tree like taxonomy
type category struct {
	ID               string   `json:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"shortDescription"`
	Description      string   `json:"description"`
	Images           []string `json:"images"`
}

// product price could vary from state to state
type product struct {
	ID               string   `json:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"shortDescription"`
	Description      string   `json:"description"`
	Images           []string `json:"images"`
	Brand            brand    `json:"brand"`
	Category         category `json:"category"`
}

var productList = []product{
	{ID: "1", Title: "UB Export", ShortDescription: "", Description: "", Images: []string{"", ""}, Brand: brands[0], Category: categories[0]},
	{ID: "2", Title: "Kingfisher Strong", ShortDescription: "", Description: "", Images: []string{"", ""}, Brand: brands[0], Category: categories[0]},
	{ID: "3", Title: "Kingfisher Premium", ShortDescription: "", Description: "", Images: []string{"", ""}, Brand: brands[0], Category: categories[0]},
}

var categories = []category{
	{ID: "1", Title: "Beer - Alcoholic Beverage", ShortDescription: "", Description: "", Images: []string{"", ""}},
}

var brands = []brand{
	{ID: "1", Title: "Kingfisher", ShortDescription: "", Description: "", Images: []string{"", ""}},
}
