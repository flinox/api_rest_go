package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/flinox/api_rest_go/models"
	"github.com/flinox/api_rest_go/utils"
	"github.com/gorilla/mux"
)

var (
	products []models.Product
)

// GetAllProducts Get all products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	defer utils.TimeTrack(time.Now(), "GetAllProducts")
	json.NewEncoder(w).Encode(products)
}

// GetProduct Get specific product by id
func GetProduct(w http.ResponseWriter, r *http.Request) {
	defer utils.TimeTrack(time.Now(), "GetProduct")
	params := mux.Vars(r)
	for _, product := range products {
		if product.ID == params["id"] {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Product{})
}

// CreateProduct Create a product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	defer utils.TimeTrack(time.Now(), "CreateProduct")
	params := mux.Vars(r)
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = params["id"]
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct Delete product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	defer utils.TimeTrack(time.Now(), "DeleteProduct")
	params := mux.Vars(r)
	for index, product := range products {
		if product.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			w.WriteHeader(200)
			break
		}
	}
	w.WriteHeader(404)
}
