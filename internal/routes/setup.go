package routes

import (
	"net/http"

	"ecommerce/internal/services"
)

func SetupRoutes(mux *http.ServeMux, productSvc *services.ProductService) {
	product := NewProductHandler(productSvc)
	mux.HandleFunc("GET /products", product.GetAll)
	mux.HandleFunc("GET /products/{id}", product.GetByID)
	mux.HandleFunc("POST /products", product.Create)
}
