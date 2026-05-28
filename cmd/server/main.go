package cmd

import (
	"context"
	"ecommerce/internal/config"
	"ecommerce/internal/database"
	"ecommerce/internal/database/repository"
	"ecommerce/internal/middleware"
	"ecommerce/internal/routes"
	"ecommerce/internal/services"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		logFile = os.Stderr
	}
	defer logFile.Close()

	handler := slog.NewJSONHandler(io.MultiWriter(logFile, os.Stdout), &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Info("🪵 Logger initialized", "output", "app.log + console")

	// 📦 2️⃣ Load config
	cnf := config.Get()
	addr := fmt.Sprintf(":%d", cnf.Port)

	// 🔌 3️⃣ Connect to Database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(ctx, cnf.DatabaseURL, int32(cnf.DatabaseMaxConns))
	if err != nil {
		slog.Error("❌ DB connection failed", "error", err)
		return
	}
	defer db.Close()
	slog.Info("✅ Database connection established")

	// 🔄 4️⃣ Run Migrations
	if err := database.RunMigrations(cnf.DatabaseURL, "./migrations"); err != nil {
		slog.Error("❌ Migration failed", "error", err)
		return
	}

	// 🗄️ 5️⃣ Wire Repositories & Services
	productRepo := repository.NewProductRepository(db)
	productSvc := services.NewProductService(
		productRepo, productRepo, productRepo, productRepo, productRepo,
	)

	// 🛣️ 6️⃣ Setup Routes
	mux := http.NewServeMux()
	routes.SetupRoutes(mux, productSvc)

	// 🔗 7️⃣ Apply Middleware
	mw := middleware.New()
	mw.Use(middleware.CORS, middleware.Logger)
	handlerChain := mw.Chain(mux)

	// 🚀 8️⃣ Start HTTP Server
	server := &http.Server{
		Addr:         addr,
		Handler:      handlerChain,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	slog.Info("🚀 Server starting", "address", addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("❌ Server crashed", "error", err)
		}
	}()

	// 🛑 9️⃣ Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("🛑 Shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("❌ Forced shutdown", "error", err)
	}
	slog.Info("✨ Server exited")
}
