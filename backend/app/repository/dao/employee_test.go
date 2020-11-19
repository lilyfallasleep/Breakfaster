package dao

import (
	exc "breakfaster/pkg/exception"
	"breakfaster/repository/model"
	"database/sql"
	"io/ioutil"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	dblog "gorm.io/gorm/logger"
)

type EmployeeSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository              EmployeeRepository
	selectIDEmployeeColumns []string
	empID                   string
	lineUID                 string
}

func TestEmployee(t *testing.T) {
	suite.Run(t, new(EmployeeSuite))
}

// This will run before before the tests in the suite
func (s *EmployeeSuite) SetupSuite() {
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

	s.repository = NewEmployeeRepository(s.DB, dummyConfig)

	s.selectIDEmployeeColumns = []string{"emp_id", "line_uid"}
}

func (s *EmployeeSuite) SetupTest() {
	s.empID = "LW99999"
	s.lineUID = "U6664ceab1f4466b30827d936cee888e6"
}

func (s *EmployeeSuite) AfterTest(suiteName, testName string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *EmployeeSuite) TestGetEmployeeID() {
	s.mock.ExpectQuery(
		"^SELECT `emp_id`,`line_uid` FROM `employees` WHERE line_uid = \\? ORDER BY `employees`.`emp_id` LIMIT 1").
		WithArgs(s.lineUID).
		WillReturnRows(sqlmock.NewRows(s.selectIDEmployeeColumns).
			AddRow(s.empID, s.lineUID),
		)

	empID, err := s.repository.GetEmpID(s.lineUID)
	require.NoError(s.T(), err)
	assert.EqualValues(s.T(), s.empID, empID)
}

func (s *EmployeeSuite) TestGetLineUID() {
	s.mock.ExpectQuery(
		"^SELECT `emp_id`,`line_uid` FROM `employees` WHERE emp_id = \\? ORDER BY `employees`.`emp_id` LIMIT 1").
		WithArgs(s.empID).
		WillReturnRows(sqlmock.NewRows(s.selectIDEmployeeColumns).
			AddRow(s.empID, s.lineUID),
		)

	lineUID, err := s.repository.GetLineUID(s.empID)
	require.NoError(s.T(), err)
	assert.EqualValues(s.T(), s.lineUID, lineUID)
}

func (s *EmployeeSuite) TestGetEmployeeIDNotFound() {
	s.mock.ExpectQuery(
		"^SELECT `emp_id`,`line_uid` FROM `employees` WHERE line_uid = \\? ORDER BY `employees`.`emp_id` LIMIT 1").
		WithArgs(s.lineUID).
		WillReturnRows(sqlmock.NewRows(s.selectIDEmployeeColumns))

	_, err := s.repository.GetEmpID(s.lineUID)
	assert.EqualError(s.T(), err, exc.ErrEmployeeNotFound.Error())
}

func (s *EmployeeSuite) TestGetLineUIDNotFound() {
	s.mock.ExpectQuery(
		"^SELECT `emp_id`,`line_uid` FROM `employees` WHERE emp_id = \\? ORDER BY `employees`.`emp_id` LIMIT 1").
		WithArgs(s.empID).
		WillReturnRows(sqlmock.NewRows(s.selectIDEmployeeColumns))

	_, err := s.repository.GetLineUID(s.empID)
	assert.EqualError(s.T(), err, exc.ErrEmployeeNotFound.Error())
}

func (s *EmployeeSuite) TestInsertNewEmployee() {
	createdAt := time.Now().Unix()
	updatedAt := createdAt
	employee := &model.Employee{
		EmpID:   s.empID,
		LineUID: s.lineUID,
	}
	s.mock.ExpectExec("^INSERT INTO `employees` \\(`emp_id`,`line_uid`,`updated_at`,`created_at`\\) VALUES (.+) ON DUPLICATE KEY UPDATE `line_uid`=VALUES\\(`line_uid`\\),`emp_id`=VALUES\\(`emp_id`\\),`updated_at`=VALUES\\(`updated_at`\\),`created_at`=VALUES\\(`created_at`\\)$").
		WithArgs(employee.EmpID, employee.LineUID, createdAt, updatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1)) // inserted id = 1, 1 affected row
	err := s.repository.UpsertEmployeeByIDs(employee)
	require.NoError(s.T(), err)
}
