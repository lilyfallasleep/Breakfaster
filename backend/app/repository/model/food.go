package model

import "time"

// Food is food model type
type Food struct {
	ID             int       `gorm:"primaryKey"`
	FoodName       string    `gorm:"type:varchar(128)"`
	FoodSupplier   string    `gorm:"type:varchar(128)"`
	PicURL         string    `gorm:"type:varchar(512)"`
	SupplyDatetime time.Time `gorm:"index"`
	UpdatedAt      int64     `gorm:"autoUpdateTime"`
	CreatedAt      int64     `gorm:"autoCreateTime"`
}
