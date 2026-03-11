package dto

type GetStockByProductIdReq struct{
	ProductID int `uri:"productId" binding:"required"`
}
type GetStockByProductIdRes struct {
	ProductID, AvailableStock,ReservedStock int
}
