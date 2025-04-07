package router

import (
	"net/http"

	"github.com/ecomz/backend/product-service/internal/handler"
	"github.com/gorilla/mux"
)

func NewRouter(categoryHandler *handler.CategoryHandler, productHandler *handler.ProductHandler) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	product := api.PathPrefix("/products").Subrouter()

	product.HandleFunc("", productHandler.GetAllProducts).Methods(http.MethodGet)
	product.HandleFunc("", productHandler.CreateProduct).Methods(http.MethodPost)
	product.HandleFunc("/{id}", productHandler.GetProductByID).Methods(http.MethodGet)
	product.HandleFunc("/{id}", productHandler.UpdateProduct).Methods(http.MethodPut)
	product.HandleFunc("/{id}", productHandler.DeleteProduct).Methods(http.MethodDelete)

	product.HandleFunc("/categories", categoryHandler.GetAllCategories).Methods(http.MethodGet)
	product.HandleFunc("/categories", categoryHandler.CreateCategory).Methods(http.MethodPost)
	product.HandleFunc("/categories/{id}", categoryHandler.GetCategoryByID).Methods(http.MethodGet)
	product.HandleFunc("/categories/{id}", categoryHandler.UpdateCategory).Methods(http.MethodPut)
	product.HandleFunc("/categories/{id}", categoryHandler.DeleteCategory).Methods(http.MethodDelete)

	return r
}
