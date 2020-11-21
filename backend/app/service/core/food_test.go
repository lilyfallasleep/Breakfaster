package core

import (
	"breakfaster/config"
	"breakfaster/mocks/mock_cache"
	"breakfaster/mocks/mock_dao"
	exc "breakfaster/pkg/exception"
	rs "breakfaster/repository/schema"
	ss "breakfaster/service/schema"
	"io/ioutil"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FoodSuite struct {
	suite.Suite
	mockCtrl    *gomock.Controller
	dummyConfig *config.Config
}

func TestFood(t *testing.T) {
	suite.Run(t, new(FoodSuite))
}

func (s *FoodSuite) SetupSuite() {
	// discard all log outputs
	log.SetOutput(ioutil.Discard)

	s.mockCtrl = gomock.NewController(s.T())
	s.dummyConfig = NewDummyConfig()
}

func (s *FoodSuite) TearDownSuite() {
	s.mockCtrl.Finish()
}

func GetTestFoodData(startDate, endDate string, start, end time.Time) ([]rs.SelectFood, ss.NestedFood) {
	returnSelectFood := []rs.SelectFood{
		rs.SelectFood{
			ID:             1,
			FoodName:       "apple",
			FoodSupplier:   "ming",
			PicURL:         "pic.com",
			SupplyDatetime: start,
		},
		rs.SelectFood{
			ID:             2,
			FoodName:       "orange",
			FoodSupplier:   "ming",
			PicURL:         "pic2.com",
			SupplyDatetime: start,
		},
		rs.SelectFood{
			ID:             3,
			FoodName:       "banana",
			FoodSupplier:   "ming",
			PicURL:         "pic3.com",
			SupplyDatetime: end,
		},
	}
	finalNestedFood := ss.NestedFood{
		startDate: []ss.JSONFood{
			ss.JSONFood{
				ID:       returnSelectFood[0].ID,
				Name:     returnSelectFood[0].FoodName,
				Supplier: returnSelectFood[0].FoodSupplier,
				PicURL:   returnSelectFood[0].PicURL,
			},
			ss.JSONFood{
				ID:       returnSelectFood[1].ID,
				Name:     returnSelectFood[1].FoodName,
				Supplier: returnSelectFood[1].FoodSupplier,
				PicURL:   returnSelectFood[1].PicURL,
			},
		},
		endDate: []ss.JSONFood{
			ss.JSONFood{
				ID:       returnSelectFood[2].ID,
				Name:     returnSelectFood[2].FoodName,
				Supplier: returnSelectFood[2].FoodSupplier,
				PicURL:   returnSelectFood[2].PicURL,
			},
		},
	}
	return returnSelectFood, finalNestedFood
}

func (s *FoodSuite) TestGetAllFoodInTimeIntervalCacheMissed() {
	startDate, endDate, start, end := GetTestDateTimeInterval()
	inputNestedFood := make(ss.NestedFood)
	returnSelectFood, finalNestedFood := GetTestFoodData(startDate, endDate, start, end)

	mockRedisCache := mock_cache.NewMockRedisCache(s.mockCtrl)
	mockFoodRepository := mock_dao.NewMockFoodRepository(s.mockCtrl)

	mockRedisCache.EXPECT().
		Get(getFoodsCacheKey(startDate, endDate), &inputNestedFood).
		Return(false, nil)
	mockFoodRepository.EXPECT().
		GetFoodAll(start, end).
		Return(&returnSelectFood, nil).Times(1)
	mockRedisCache.EXPECT().
		Set(getFoodsCacheKey(startDate, endDate), &finalNestedFood).
		Return(nil)

	foodSvc := NewFoodService(mockFoodRepository, mockRedisCache, s.dummyConfig)
	res, err := foodSvc.GetFoodAll(startDate, endDate)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&finalNestedFood, res))
}

func (s *FoodSuite) TestGetAllFoodInTimeIntervalNotFound() {
	startDate, endDate, start, end := GetTestDateTimeInterval()
	inputNestedFood := make(ss.NestedFood)

	mockRedisCache := mock_cache.NewMockRedisCache(s.mockCtrl)
	mockFoodRepository := mock_dao.NewMockFoodRepository(s.mockCtrl)

	mockRedisCache.EXPECT().
		Get(getFoodsCacheKey(startDate, endDate), &inputNestedFood).
		Return(false, nil)
	mockFoodRepository.EXPECT().
		GetFoodAll(start, end).
		Return(&[]rs.SelectFood{}, exc.ErrFoodNotFound).Times(1)

	foodSvc := NewFoodService(mockFoodRepository, mockRedisCache, s.dummyConfig)
	_, err := foodSvc.GetFoodAll(startDate, endDate)
	assert.EqualError(s.T(), err, exc.ErrFoodNotFound.Error())
}

func (s *FoodSuite) TestGetAllFoodInTimeIntervalCached() {
	startDate, endDate, _, _ := GetTestDateTimeInterval()
	inputNestedFood := make(ss.NestedFood)

	mockRedisCache := mock_cache.NewMockRedisCache(s.mockCtrl)
	mockFoodRepository := mock_dao.NewMockFoodRepository(s.mockCtrl)

	mockRedisCache.EXPECT().
		Get(getFoodsCacheKey(startDate, endDate), &inputNestedFood).
		Return(true, nil)

	foodSvc := NewFoodService(mockFoodRepository, mockRedisCache, s.dummyConfig)
	_, err := foodSvc.GetFoodAll(startDate, endDate)
	require.NoError(s.T(), err)
}
