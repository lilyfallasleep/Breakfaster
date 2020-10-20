package model

import (
	"time"
)

// Order is order model type
type Order struct {
	FoodID        int
	EmployeeEmpID string    `gorm:"primaryKey"`
	Food          Food      `gorm:"foreignKey:FoodID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Employee      Employee  `gorm:"foreignKey:EmployeeEmpID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Date          time.Time `gorm:"primaryKey"`
	Pick          bool      `gorm:"default:false"`
	PickUpAt      int64
	UpdatedAt     int64 `gorm:"autoUpdateTime"`
	CreatedAt     int64 `gorm:"autoCreateTime"`
}
