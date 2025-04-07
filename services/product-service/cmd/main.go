package main

import (
	"net/http"

	"github.com/ecomz/backend/libs/config"
	"github.com/ecomz/backend/libs/db"
	"github.com/ecomz/backend/libs/logger"
	"github.com/ecomz/backend/product-service/cmd/router"
	"github.com/ecomz/backend/product-service/internal/handler"
	"github.com/ecomz/backend/product-service/internal/repository"
	"github.com/ecomz/backend/product-service/internal/service"
	"go.uber.org/zap"
)

func main() {
	zapLogger := logger.NewZapLogger()
	cfg := config.LoadConfigFromFile("./cmd", "config", "yml")

	dbConn, err := db.NewConnectionManager(cfg.Database)
	if err != nil {
		zapLogger.Fatal("Failed to create database connection", zap.Error(err))
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			zapLogger.Fatal("Failed to close database connection", zap.Error(err))
		}
	}()

	categoryRepository := repository.NewCategoryRepository(dbConn.GetDB())
	categoryService := service.NewCategoryService(zapLogger, categoryRepository)
	categoryHandler := handler.NewCategoryHandler(zapLogger, categoryService)

	productRepository := repository.NewProductRepository(dbConn.GetDB())
	productService := service.NewProductService(zapLogger, productRepository)
	productHandler := handler.NewProductHandler(zapLogger, productService)

	r := router.NewRouter(categoryHandler, productHandler)

	serverAddress := ":" + cfg.App.Port

	err = http.ListenAndServe(serverAddress, r)
	if err != nil {
		zapLogger.Fatal("Failed to start server", zap.Error(err))
		return
	}

	zapLogger.Info("Server started", zap.String("address", serverAddress))
}
