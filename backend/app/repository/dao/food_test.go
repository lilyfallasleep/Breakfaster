package dao

import (
	exc "breakfaster/pkg/exception"
	"breakfaster/repository/schema"
	"breakfaster/service/constant"
	"database/sql"
	"io/ioutil"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	dblog "gorm.io/gorm/logger"
)

type FoodSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository        FoodRepository
	selectFoodColumns []string
	start             time.Time
	end               time.Time
}

func TestFoodRepositoryImpl(t *testing.T) {
	suite.Run(t, new(FoodSuite))
}

// This will run before before the tests in the suite
func (s *FoodSuite) SetupSuite() {
	// discard all log outputs
	log.SetOutput(ioutil.Discard)

	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dummyConfig := NewDummyConfig()

	s.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: dblog.Default.LogMode(logger.Silent),
	})
	require.NoError(s.T(), err)

	s.repository = NewFoodRepository(s.DB, dummyConfig)

	s.selectFoodColumns = []string{"id", "food_name", "food_supplier", "pic_url", "supply_datetime"}
}

func (s *FoodSuite) SetupTest() {
	start, _ := time.ParseInLocation(constant.DateFormat, "2020-10-10", time.Local)
	end, _ := time.ParseInLocation(constant.DateFormat, "2020-10-11", time.Local)

	s.start = start
	s.end = end
}

// This will run after each test finishes
func (s *FoodSuite) AfterTest(suiteName, testName string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *FoodSuite) TestGetAllFoodInTimeInterval() {
	testFoods := []schema.SelectFood{
		schema.SelectFood{
			ID:             1,
			FoodName:       "apple",
			FoodSupplier:   "ming",
			PicURL:         "https://pic.com",
			SupplyDatetime: s.start,
		},
		schema.SelectFood{
			ID:             2,
			FoodName:       "apple",
			FoodSupplier:   "ming",
			PicURL:         "https://pic.com",
			SupplyDatetime: s.start,
		},
		schema.SelectFood{
			ID:             3,
			FoodName:       "banana",
			FoodSupplier:   "ming",
			PicURL:         "https://pic.com",
			SupplyDatetime: s.end,
		},
	}

	validFoods := []schema.SelectFood{testFoods[0], testFoods[1], testFoods[2]}

	s.mock.ExpectQuery(
		"^SELECT `id`,`food_name`,`food_supplier`,`pic_url`,`supply_datetime` FROM `foods` WHERE supply_datetime BETWEEN \\? AND \\?$").
		WithArgs(s.start, s.end).
		WillReturnRows(sqlmock.NewRows(s.selectFoodColumns).
			AddRow(testFoods[0].ID, testFoods[0].FoodName, testFoods[0].FoodSupplier, testFoods[0].PicURL, testFoods[0].SupplyDatetime).
			AddRow(testFoods[1].ID, testFoods[1].FoodName, testFoods[1].FoodSupplier, testFoods[1].PicURL, testFoods[1].SupplyDatetime).
			AddRow(testFoods[2].ID, testFoods[2].FoodName, testFoods[2].FoodSupplier, testFoods[2].PicURL, testFoods[2].SupplyDatetime),
		)

	res, err := s.repository.GetFoodAll(s.start, s.end)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&validFoods, res), "should return all valid foods")
}

func (s *FoodSuite) TestFoodNotFound() {

	s.mock.ExpectQuery(
		"^SELECT `id`,`food_name`,`food_supplier`,`pic_url`,`supply_datetime` FROM `foods` WHERE supply_datetime BETWEEN \\? AND \\?$").
		WithArgs(s.start, s.end).
		WillReturnRows(sqlmock.NewRows(s.selectFoodColumns)) // return nothing

	_, err := s.repository.GetFoodAll(s.start, s.end)
	assert.EqualError(s.T(), err, exc.ErrFoodNotFound.Error())
}
