package routes

import (
	"net/http"

	"ecommerce/internal/services"
)

func SetupRoutes(mux *http.ServeMux, svc *services.Service) {
	handler := NewProductHandler(svc)

	mux.HandleFunc("GET /products", handler.GetAll)
	mux.HandleFunc("GET /products/{id}", handler.GetByID)
	mux.HandleFunc("POST /products", handler.Create)
}
