package usecase

import (
	"context"
	"go-inventory-service/internal/application/dto"
	"go-inventory-service/internal/domain/repository"
)

type InventoryUsecase struct {
	repo repository.InventoryRepo
}

func NewInventoryUsecase(r repository.InventoryRepo) *InventoryUsecase {
	return &InventoryUsecase{repo: r}
}

func (uc *InventoryUsecase) GetStockByProductId(ctx context.Context, req dto.GetStockByProductIdReq) (*dto.GetStockByProductIdRes, error) {
	stock,err := uc.repo.GetStockByProductId(ctx, req.ProductID)
	if err != nil {
		return nil, err
	}
	return &dto.GetStockByProductIdRes{
		ProductID: stock.ProductID,
		AvailableStock: int(stock.AvailableStock),
		ReservedStock: int(stock.ReservedStock),
	}, nil
}