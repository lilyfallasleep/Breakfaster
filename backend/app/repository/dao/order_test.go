package dao

import (
	exc "breakfaster/pkg/exception"
	"breakfaster/repository/model"
	"breakfaster/repository/schema"
	"breakfaster/service/constant"
	"database/sql"
	"io/ioutil"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/DATA-DOG/go-sqlmock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//dblog "gorm.io/gorm/logger"
)

type OrderSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository                          OrderRepository
	selectOrderColumns                  []string
	selectOrderWithEmployeeEmpIDColumns []string
}

func TestOrder(t *testing.T) {
	suite.Run(t, new(OrderSuite))
}

// This will run before before the tests in the suite
func (s *OrderSuite) SetupSuite() {
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
		//Logger: dblog.Default.LogMode(logger.Silent),
	})
	require.NoError(s.T(), err)

	s.repository = NewOrderRepository(s.DB, dummyConfig)

	s.selectOrderColumns = []string{"food_name", "date"}
	s.selectOrderWithEmployeeEmpIDColumns = []string{"food_name", "employee_emp_id", "date", "pick"}
}

func (s *OrderSuite) AfterTest(suiteName, testName string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *OrderSuite) TestCreateOrders() {
	date, _ := time.ParseInLocation(constant.DateFormat, "2020-10-10", time.Local)
	orders := &[]model.Order{
		model.Order{
			FoodID:        1,
			EmployeeEmpID: "LW99999",
			Date:          date,
		},
	}
	s.mock.ExpectExec("^INSERT INTO `orders` \\(`food_id`,`employee_emp_id`,`date`,`updated_at`,`created_at`\\) VALUES (.+) ON DUPLICATE KEY UPDATE `food_id`=VALUES\\(`food_id`\\),`updated_at`=VALUES\\(`updated_at`\\),`created_at`=VALUES\\(`created_at`\\)$").
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := s.repository.CreateOrders(orders)
	require.NoError(s.T(), err)
}

func (s *OrderSuite) TestDeleteOrdersByLineUID() {
	start, _ := time.ParseInLocation(constant.DateFormat, "2020-10-10", time.Local)
	end, _ := time.ParseInLocation(constant.DateFormat, "2020-10-10", time.Local)
	lineUID := "U6664ceab1f4466b30827d936cee888e6"
	s.mock.ExpectExec("^DELETE FROM `orders` WHERE \\(date BETWEEN \\? AND \\?\\) AND employee_emp_id = \\(SELECT emp_id FROM `employees` WHERE line_uid = \\?\\)$").
		WithArgs(start, end, lineUID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := s.repository.DeleteOrdersByLineUID(lineUID, start, end)
	require.NoError(s.T(), err)
}

func (s *OrderSuite) TestGetOrdersByLineUID() {
	start, _ := time.ParseInLocation(constant.DateFormat, "2020-10-10", time.Local)
	end, _ := time.ParseInLocation(constant.DateFormat, "2020-10-10", time.Local)
	lineUID := "U6664ceab1f4466b30827d936cee888e6"
	orders := []schema.SelectOrder{
		schema.SelectOrder{
			FoodName: "apple",
			Date:     start,
		},
		schema.SelectOrder{
			FoodName: "orange",
			Date:     end,
		},
	}
	s.mock.ExpectQuery(
		"^SELECT orders.date,foods.food_name FROM `orders` left join foods on foods.id = orders.food_id left join employees on employees.emp_id = orders.employee_emp_id WHERE \\(date BETWEEN \\? AND \\?\\) AND employees.line_uid = \\?$").
		WithArgs(start, end, lineUID).
		WillReturnRows(sqlmock.NewRows(s.selectOrderColumns).
			AddRow(orders[0].FoodName, orders[0].Date).
			AddRow(orders[1].FoodName, orders[1].Date),
		)

	res, err := s.repository.GetOrdersByLineUID(lineUID, start, end)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&orders, res))
}

func (s *OrderSuite) TestGetOrderByEmpID() {
	date, _ := time.ParseInLocation(constant.DateFormat, "2020-10-10", time.Local)
	empID := "LW99999"
	order := schema.SelectOrderWithEmployeeEmpID{
		FoodName:      "apple",
		EmployeeEmpID: empID,
		Date:          date,
		Pick:          false,
	}

	s.mock.ExpectQuery(
		"^SELECT orders.date,orders.employee_emp_id,orders.pick,foods.food_name FROM `orders` left join foods on foods.id = orders.food_id WHERE date = \\? AND \\(orders.employee_emp_id = \\?\\) ORDER BY `orders`.`food_id` LIMIT 1$").
		WithArgs(date, empID).
		WillReturnRows(sqlmock.NewRows(s.selectOrderWithEmployeeEmpIDColumns).
			AddRow(order.FoodName, order.EmployeeEmpID, order.Date, order.Pick),
		)

	res, err := s.repository.GetOrderByEmpID(empID, date)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&order, res))
}

func (s *OrderSuite) TestGetOrderByEmpIDNotFound() {
	date, _ := time.ParseInLocation(constant.DateFormat, "2020-10-10", time.Local)
	empID := "LW99999"
	s.mock.ExpectQuery(
		"^SELECT orders.date,orders.employee_emp_id,orders.pick,foods.food_name FROM `orders` left join foods on foods.id = orders.food_id WHERE date = \\? AND \\(orders.employee_emp_id = \\?\\) ORDER BY `orders`.`food_id` LIMIT 1$").
		WithArgs(date, empID).
		WillReturnRows(sqlmock.NewRows(s.selectOrderWithEmployeeEmpIDColumns))

	_, err := s.repository.GetOrderByEmpID(empID, date)
	assert.EqualError(s.T(), err, exc.ErrOrderNotFound.Error())
}

func (s *OrderSuite) TestGetOrderByAccessCardNumber() {
	date, _ := time.ParseInLocation(constant.DateFormat, "2020-10-10", time.Local)
	empID := "LW99999"
	accessCardNbr := "1234567890"
	order := schema.SelectOrderWithEmployeeEmpID{
		FoodName:      "apple",
		EmployeeEmpID: empID,
		Date:          date,
		Pick:          false,
	}

	s.mock.ExpectQuery(
		"^SELECT orders.date,orders.employee_emp_id,orders.pick,foods.food_name FROM `orders` left join foods on foods.id = orders.food_id left join employees on employees.emp_id = orders.employee_emp_id WHERE date = \\? AND employees.access_card_nbr = \\? ORDER BY `orders`.`food_id` LIMIT 1$").
		WithArgs(date, accessCardNbr).
		WillReturnRows(sqlmock.NewRows(s.selectOrderWithEmployeeEmpIDColumns).
			AddRow(order.FoodName, order.EmployeeEmpID, order.Date, order.Pick),
		)

	res, err := s.repository.GetOrderByAccessCardNumber(accessCardNbr, date)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&order, res))
}

func (s *OrderSuite) TestGetOrderByAccessCardNumberNotFound() {
	date, _ := time.ParseInLocation(constant.DateFormat, "2020-10-10", time.Local)
	accessCardNbr := "1234567890"

	s.mock.ExpectQuery(
		"^SELECT orders.date,orders.employee_emp_id,orders.pick,foods.food_name FROM `orders` left join foods on foods.id = orders.food_id left join employees on employees.emp_id = orders.employee_emp_id WHERE date = \\? AND employees.access_card_nbr = \\? ORDER BY `orders`.`food_id` LIMIT 1$").
		WithArgs(date, accessCardNbr).
		WillReturnRows(sqlmock.NewRows(s.selectOrderWithEmployeeEmpIDColumns))

	_, err := s.repository.GetOrderByAccessCardNumber(accessCardNbr, date)
	assert.EqualError(s.T(), err, exc.ErrOrderNotFound.Error())
}

func (s *OrderSuite) TestUpdateOrderStatus() {
	date, _ := time.ParseInLocation(constant.DateFormat, "2020-10-10", time.Local)
	empID := "LW99999"
	pick := true
	pickUpAt := time.Now().Unix()

	s.mock.ExpectExec("^UPDATE `orders` SET `pick`=\\?,`pick_up_at`=\\?,`updated_at`=\\? WHERE employee_emp_id = \\? AND date = \\?$").
		WithArgs(pick, pickUpAt, pickUpAt, empID, date).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := s.repository.UpdateOrderStatus(empID, date, pick, pickUpAt)
	require.NoError(s.T(), err)
}
