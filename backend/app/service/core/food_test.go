package core

import (
	"breakfaster/config"
	"breakfaster/infrastructure/cache"
	"breakfaster/repository/dao"
	ss "breakfaster/service/schema"
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

func Test_getFoodsCacheKey(t *testing.T) {
	type args struct {
		startDate string
		endDate   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFoodsCacheKey(tt.args.startDate, tt.args.endDate); got != tt.want {
				t.Errorf("getFoodsCacheKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoodServiceImpl_GetFoodAll(t *testing.T) {
	type fields struct {
		repository dao.FoodRepository
		cache      cache.RedisCache
		logger     *log.Entry
	}
	type args struct {
		startDate string
		endDate   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ss.NestedFood
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &FoodServiceImpl{
				repository: tt.fields.repository,
				cache:      tt.fields.cache,
				logger:     tt.fields.logger,
			}
			got, err := svc.GetFoodAll(tt.args.startDate, tt.args.endDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("FoodServiceImpl.GetFoodAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoodServiceImpl.GetFoodAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFoodService(t *testing.T) {
	type args struct {
		repository dao.FoodRepository
		cache      cache.RedisCache
		config     *config.Config
	}
	tests := []struct {
		name string
		args args
		want FoodService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFoodService(tt.args.repository, tt.args.cache, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFoodService() = %v, want %v", got, tt.want)
			}
		})
	}
}
