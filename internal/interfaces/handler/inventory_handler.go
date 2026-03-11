package handler

import (
	"go-inventory-service/internal/application/dto"
	"go-inventory-service/internal/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	usecase *usecase.InventoryUsecase
}

func NewInventoryHandler(uc *usecase.InventoryUsecase) *InventoryHandler {
	return &InventoryHandler{usecase: uc}
}

func (h *InventoryHandler) GetStockByProductId(c *gin.Context) {
	var req dto.GetStockByProductIdReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	res,err := h.usecase.GetStockByProductId(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}
	c.JSON(http.StatusOK, res)
}

