package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ecomz/backend/auth-service/internal/handler"
	"github.com/ecomz/backend/auth-service/internal/repository"
	"github.com/ecomz/backend/auth-service/internal/router"
	"github.com/ecomz/backend/auth-service/internal/service"
	"github.com/ecomz/backend/libs/config"
	"github.com/ecomz/backend/libs/db"
	"github.com/ecomz/backend/libs/utils"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
	defer logger.Sync()
}

func main() {
	cfg := config.LoadConfigFromFile("./cmd", "config", "yml")
	dbConn, err := db.NewConnectionManager(cfg.Database)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Printf("error closing database connection: %v", err)
		}
	}()

	redisAddr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)

	pool, err := utils.CreateRedisPool(redisAddr, cfg.Redis.Password, cfg.Redis.MaxIdle, logger)
	if err != nil {
		logger.Fatal("Failed to create Redis pool", zap.Error(err))
	}
	utils.NewCacheService(pool, logger)
	logger.Info("Successfully connected to Redis")

	sqlDB := dbConn.GetDB()

	userRepo := repository.NewUserRepository(sqlDB, pool)
	roleRepo := repository.NewRoleRepository(sqlDB)

	userService := service.NewUserService(logger, cfg, userRepo, roleRepo)
	roleService := service.NewRoleService(logger, roleRepo)

	userHandler := handler.NewUserHandler(userService)
	roleHandler := handler.NewRoleHandler(roleService)

	r := router.NewRouter(userHandler, roleHandler)

	serverAddress := ":" + cfg.App.Port
	log.Printf("Starting server on %s", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))
}
