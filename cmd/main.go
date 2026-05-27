package cmd

import (
	"ecommerce/internal/middleware"
	"ecommerce/internal/utils"
	"log"
	"net/http"
)

type Product struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Id: 1, Name: "Laptop", Price: 99999},
		{Id: 2, Name: "Wireless Mouse", Price: 2500},
		{Id: 3, Name: "Mechanical Keyboard", Price: 7500},
	}

	utils.WriteJSON(w, 200, products)
}
func Run() {
	mux := http.NewServeMux()
	addr := ":8000"

	mux.HandleFunc("GET /products", productsHandler)

	// middleware
	middleWare := middleware.New()
	middleWare.Use(middleware.CORS)
	handler := middleWare.Chain(mux)

	log.Println("Server running on", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
