package http

import (
	"go-inventory-service/internal/interfaces/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(inventoryHandler *handler.InventoryHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/api/products/:productId/stock", inventoryHandler.GetStockByProductId)

	return router
}
