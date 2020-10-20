package core

import (
	"breakfaster/infrastructure/cache"
	exc "breakfaster/pkg/exception"
	"breakfaster/repository/dao"
	rs "breakfaster/repository/schema"
	"breakfaster/service/constant"
	ss "breakfaster/service/schema"
	"fmt"
	"time"
)

// FoodService provides methods for manipulating foods
type FoodService struct {
	repository *dao.FoodRepository
	cache      cache.GeneralCache
}

const (
	foodsCacheKeyBase string = "foods-cache-key"
)

func getFoodsCacheKey(startDate, endDate string) string {
	return fmt.Sprintf("%s-%s-%s", foodsCacheKeyBase, startDate, endDate)
}

// GetFoodAll returns all foods between the given date (includingly)
func (svc *FoodService) GetFoodAll(startDate, endDate string) (*ss.NestedFood, error) {
	var err error
	var start, end time.Time
	var foods *[]rs.SelectFood

	start, err = time.ParseInLocation(constant.DateFormat, startDate, time.Local)
	if err != nil {
		return &ss.NestedFood{}, exc.ErrDateFormat
	}
	end, err = time.ParseInLocation(constant.DateFormat, endDate, time.Local)
	if err != nil {
		return &ss.NestedFood{}, exc.ErrDateFormat
	}

	foodsCacheKey := getFoodsCacheKey(startDate, endDate)
	if nestedFood, found := svc.cache.Get(foodsCacheKey); found {
		return nestedFood.(*ss.NestedFood), nil
	}

	foods, err = svc.repository.GetFoodAll(start, end)
	if err != nil {
		return &ss.NestedFood{}, err
	}

	nestedFood := make(ss.NestedFood)
	for _, food := range *foods {
		datetime := food.SupplyDatetime.Format(constant.DateFormat)
		nestedFood[datetime] = append(nestedFood[datetime], ss.JSONFood{
			ID:       food.ID,
			Name:     food.FoodName,
			Supplier: food.FoodSupplier,
			PicURL:   food.PicURL,
		})
	}
	svc.cache.Set(foodsCacheKey, &nestedFood)

	return &nestedFood, nil
}

// NewFoodService is the factory for FoodService
func NewFoodService(repository *dao.FoodRepository, cache cache.GeneralCache) *FoodService {
	return &FoodService{
		repository: repository,
		cache:      cache,
	}
}
