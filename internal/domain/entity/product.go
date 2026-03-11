package entity

type Product struct {
	ID int `gorm:"primaryKey"`
	Name string `gorm:"not null;index"`
	CreatedAt int64 `gorm:"autoCreateTime;index"`
	UpdatedAt int64 `gorm:"autoUpdateTime;index"`
}