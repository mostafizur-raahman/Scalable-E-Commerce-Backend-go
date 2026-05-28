package cmd

import (
	"ecommerce/internal/config"
	"ecommerce/internal/database"
	"ecommerce/internal/middleware"
	"ecommerce/internal/routes"
	"ecommerce/internal/services"
	"log"
	"net/http"
	"strconv"
)

func Run() {
	// 1️⃣ Load configuration
	cnf := config.Get()
	addr := ":" + strconv.Itoa(cnf.Port)

	// 2️⃣ Initialize Data Layer (Repository)
	productRepo := database.NewProductRepository()

	svc := services.NewProductService(productRepo, productRepo, productRepo)

	mux := http.NewServeMux()
	routes.SetupRoutes(mux, svc)

	// Apply Middleware Chain
	mw := middleware.New()
	mw.Use(middleware.CORS, middleware.Logger)
	handler := mw.Chain(mux)

	log.Printf("🚀 Starting %s v%s on %s", cnf.AppName, cnf.Version, addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
