package entity

type InventoryLog struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	ProductID uint   `gorm:"index"`
	Quantity  int
	Source    int
	SourceID  uint
	CreatedAt int64  `gorm:"autoCreateTime;index"`
}