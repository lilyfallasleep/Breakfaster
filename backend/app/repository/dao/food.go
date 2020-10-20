package dao

import (
	exc "breakfaster/pkg/exception"
	"breakfaster/repository/model"
	"breakfaster/repository/schema"
	"log"
	"time"

	"gorm.io/gorm"
)

// FoodRepository provides operations on food model
type FoodRepository struct {
	db *gorm.DB
}

// GetFoodAll returns all foods within the given time interval
func (repo *FoodRepository) GetFoodAll(start, end time.Time) (*[]schema.SelectFood, error) {
	var foods []schema.SelectFood
	if err := repo.db.Model(&model.Food{}).Where("supply_datetime BETWEEN ? AND ?", start, end).Find(&foods).Error; err != nil {
		log.Print(err)
		return &[]schema.SelectFood{}, exc.ErrGetFood
	}
	if len(foods) == 0 {
		return &[]schema.SelectFood{}, exc.ErrFoodNotFound
	}
	return &foods, nil
}

// NewFoodRepository is the factory for FoodRepository
func NewFoodRepository(db *gorm.DB) *FoodRepository {
	return &FoodRepository{db: db}
}
