package repository_impl

import (
	"context"
	"go-inventory-service/internal/domain/entity"

	"gorm.io/gorm"
)

type InventoryRepoImpl struct {
	db *gorm.DB
}

func NewInventoryRepoImpl(db *gorm.DB) *InventoryRepoImpl{
	return &InventoryRepoImpl{db: db}
}

func (r *InventoryRepoImpl) GetStockByProductId(ctx context.Context, product_id int) (*entity.Inventory, error){
	var inventory entity.Inventory
	err := r.db.WithContext(ctx).Where("product_id=?", product_id).First(&inventory).Error
	return &inventory, err
}