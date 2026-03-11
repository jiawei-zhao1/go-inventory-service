package entity

type Inventory struct {
	ProductID       int   `gorm:"primaryKey"`
	AvailableStock  uint
	ReservedStock   uint
	CreatedAt       int64 `gorm:"autoCreateTime;index"`
	UpdatedAt       int64 `gorm:"autoUpdateTime;index"`
}