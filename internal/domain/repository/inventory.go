package repository

import (
	"context"
	"go-inventory-service/internal/domain/entity"
)

type InventoryRepo interface {
	GetStockByProductId(ctx context.Context, product_id int) (*entity.Inventory, error)
}