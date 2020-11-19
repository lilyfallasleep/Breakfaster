package dao

import (
	"breakfaster/config"
	exc "breakfaster/pkg/exception"
	"breakfaster/repository/model"
	"breakfaster/repository/schema"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// FoodRepositoryImpl provides operations on food model
type FoodRepositoryImpl struct {
	db     *gorm.DB
	logger *log.Entry
}

// GetFoodAll returns all foods within the given time interval
func (repo *FoodRepositoryImpl) GetFoodAll(start, end time.Time) (*[]schema.SelectFood, error) {
	var foods []schema.SelectFood
	rows, err := repo.db.Model(&model.Food{}).Where("supply_datetime BETWEEN ? AND ?", start, end).
		Select("id", "food_name", "food_supplier", "pic_url", "supply_datetime").Rows()
	defer rows.Close()
	if err != nil {
		repo.logger.Error(err)
		return &[]schema.SelectFood{}, exc.ErrGetFood
	}

	var food schema.SelectFood
	for rows.Next() {
		repo.db.ScanRows(rows, &food)
		foods = append(foods, food)
	}
	if len(foods) == 0 {
		return &[]schema.SelectFood{}, exc.ErrFoodNotFound
	}
	return &foods, nil
}

// NewFoodRepository is the factory for FoodRepositoryImpl
func NewFoodRepository(db *gorm.DB, config *config.Config) FoodRepository {
	return &FoodRepositoryImpl{
		db: db,
		logger: config.Logger.ContextLogger.WithFields(log.Fields{
			"type": "dao:food",
		}),
	}
}
