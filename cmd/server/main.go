package main

import (
	"go-inventory-service/internal/application/usecase"
	"go-inventory-service/internal/config"
	"go-inventory-service/internal/infrastructure/database"
	"go-inventory-service/internal/infrastructure/repository_impl"
	"go-inventory-service/internal/interfaces/handler"
	"go-inventory-service/internal/interfaces/http"
	"log"
)

func main() {
  // 1. Load config
  cfg := config.Load()

  // 2. Init database
  db := database.NewMySQL(cfg.Db)

  // 3. Init repository
  inventoryRepo := repository_impl.NewInventoryRepoImpl(db)

  // 4. Init usecase
  inventoryUsecase := usecase.NewInventoryUsecase(inventoryRepo)

  // 5. Init handler
  inventoryHandler := handler.NewInventoryHandler(inventoryUsecase)

  // 6. Init router
  router := http.NewRouter(inventoryHandler)

  // 7. Start server
  if err := router.Run(":8080"); err != nil {
    log.Fatalf("Failed to start server: %v", err)
  }
}