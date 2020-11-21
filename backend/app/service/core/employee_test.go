package core

import (
	"breakfaster/config"
	"breakfaster/infrastructure/cache"
	"breakfaster/mocks/mock_cache"
	"breakfaster/mocks/mock_dao"
	exc "breakfaster/pkg/exception"
	"breakfaster/repository/model"
	ss "breakfaster/service/schema"
	"io/ioutil"
	"testing"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type EmployeeSuite struct {
	suite.Suite
	mockCtrl    *gomock.Controller
	dummyConfig *config.Config
}

func TestEmployee(t *testing.T) {
	suite.Run(t, new(EmployeeSuite))
}

func (s *EmployeeSuite) SetupSuite() {
	// discard all log outputs
	log.SetOutput(ioutil.Discard)

	s.mockCtrl = gomock.NewController(s.T())
	s.dummyConfig = NewDummyConfig()
}

func (s *EmployeeSuite) TearDownSuite() {
	s.mockCtrl.Finish()
}

func (s *EmployeeSuite) TestGetEmployeeByLineUIDCacheMissed() {
	inputEmpID := ""
	empID := "LW99999"
	lineUID := "U6664ceab1f4466b30827d936cee888e6"

	mockRedisCache := mock_cache.NewMockRedisCache(s.mockCtrl)
	mockEmployeeRepository := mock_dao.NewMockEmployeeRepository(s.mockCtrl)

	mockRedisCache.EXPECT().
		Get(lineUID, &inputEmpID).
		Return(false, nil)
	mockEmployeeRepository.EXPECT().
		GetEmpID(lineUID).
		Return(empID, nil).Times(1)
	mockRedisCache.EXPECT().
		Set(lineUID, empID).
		Return(nil)

	employeeSvc := NewEmployeeService(mockEmployeeRepository, mockRedisCache, s.dummyConfig)
	res, err := employeeSvc.GetEmployeeByLineUID(lineUID)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&ss.JSONEmployee{
		EmpID:   empID,
		LineUID: lineUID,
	}, res))
}

func (s *EmployeeSuite) TestGetEmployeeByLineUIDNotFound() {
	inputEmpID := ""
	lineUID := "U6664ceab1f4466b30827d936cee888e6"

	mockRedisCache := mock_cache.NewMockRedisCache(s.mockCtrl)
	mockEmployeeRepository := mock_dao.NewMockEmployeeRepository(s.mockCtrl)

	mockRedisCache.EXPECT().
		Get(lineUID, &inputEmpID).
		Return(false, nil)
	mockEmployeeRepository.EXPECT().
		GetEmpID(lineUID).
		Return("", exc.ErrEmployeeNotFound).Times(1)

	employeeSvc := NewEmployeeService(mockEmployeeRepository, mockRedisCache, s.dummyConfig)
	_, err := employeeSvc.GetEmployeeByLineUID(lineUID)
	assert.EqualError(s.T(), err, exc.ErrEmployeeNotFound.Error())
}

func (s *EmployeeSuite) TestGetEmployeeByLineUIDCached() {
	inputEmpID := ""
	lineUID := "U6664ceab1f4466b30827d936cee888e6"

	mockRedisCache := mock_cache.NewMockRedisCache(s.mockCtrl)
	mockEmployeeRepository := mock_dao.NewMockEmployeeRepository(s.mockCtrl)

	mockRedisCache.EXPECT().
		Get(lineUID, &inputEmpID).
		Return(true, nil)

	employeeSvc := NewEmployeeService(mockEmployeeRepository, mockRedisCache, s.dummyConfig)
	_, err := employeeSvc.GetEmployeeByLineUID(lineUID)

	require.NoError(s.T(), err)
}

func (s *EmployeeSuite) TestGetEmployeeByEmpIDCacheMissed() {
	inputLineID := ""
	empID := "LW99999"
	lineUID := "U6664ceab1f4466b30827d936cee888e6"

	mockRedisCache := mock_cache.NewMockRedisCache(s.mockCtrl)
	mockEmployeeRepository := mock_dao.NewMockEmployeeRepository(s.mockCtrl)

	mockRedisCache.EXPECT().
		Get(empID, &inputLineID).
		Return(false, nil)
	mockEmployeeRepository.EXPECT().
		GetLineUID(empID).
		Return(lineUID, nil).Times(1)
	mockRedisCache.EXPECT().
		Set(empID, lineUID).
		Return(nil)

	employeeSvc := NewEmployeeService(mockEmployeeRepository, mockRedisCache, s.dummyConfig)
	res, err := employeeSvc.GetEmployeeByEmpID(empID)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&ss.JSONEmployee{
		EmpID:   empID,
		LineUID: lineUID,
	}, res))
}

func (s *EmployeeSuite) TestGetEmployeeByEmpIDNotFound() {
	inputLineUID := ""
	empID := "LW99999"

	mockRedisCache := mock_cache.NewMockRedisCache(s.mockCtrl)
	mockEmployeeRepository := mock_dao.NewMockEmployeeRepository(s.mockCtrl)

	mockRedisCache.EXPECT().
		Get(empID, &inputLineUID).
		Return(false, nil)
	mockEmployeeRepository.EXPECT().
		GetLineUID(empID).
		Return("", exc.ErrEmployeeNotFound).Times(1)

	employeeSvc := NewEmployeeService(mockEmployeeRepository, mockRedisCache, s.dummyConfig)
	_, err := employeeSvc.GetEmployeeByEmpID(empID)
	assert.EqualError(s.T(), err, exc.ErrEmployeeNotFound.Error())
}

func (s *EmployeeSuite) TestGetEmployeeByEmpIDCached() {
	inputLineID := ""
	empID := "LW99999"

	mockRedisCache := mock_cache.NewMockRedisCache(s.mockCtrl)
	mockEmployeeRepository := mock_dao.NewMockEmployeeRepository(s.mockCtrl)

	mockRedisCache.EXPECT().
		Get(empID, &inputLineID).
		Return(true, nil)

	employeeSvc := NewEmployeeService(mockEmployeeRepository, mockRedisCache, s.dummyConfig)
	_, err := employeeSvc.GetEmployeeByEmpID(empID)

	require.NoError(s.T(), err)
}

func (s *EmployeeSuite) TestUpsertEmployeeByIDs() {
	empID := "LW99999"
	lineUID := "U6664ceab1f4466b30827d936cee888e6"
	employee := &model.Employee{
		EmpID:   empID,
		LineUID: lineUID,
	}
	cmds := []cache.Cmd{
		cache.Cmd{
			OpType: cache.DELETE,
			Payload: cache.RedisDeletePayload{
				Key: empID,
			},
		},
		cache.Cmd{
			OpType: cache.DELETE,
			Payload: cache.RedisDeletePayload{
				Key: lineUID,
			},
		},
	}

	mockRedisCache := mock_cache.NewMockRedisCache(s.mockCtrl)
	mockEmployeeRepository := mock_dao.NewMockEmployeeRepository(s.mockCtrl)

	mockEmployeeRepository.EXPECT().
		UpsertEmployeeByIDs(employee).
		Return(nil).Times(1)
	mockRedisCache.EXPECT().
		ExecPipeLine(&cmds).
		Return(nil)

	employeeSvc := NewEmployeeService(mockEmployeeRepository, mockRedisCache, s.dummyConfig)
	err := employeeSvc.UpsertEmployeeByIDs(empID, lineUID)

	require.NoError(s.T(), err)
}

func (s *EmployeeSuite) TestUpsertEmployeeByIDsFail() {
	empID := "LW99999"
	lineUID := "U6664ceab1f4466b30827d936cee888e6"
	employee := &model.Employee{
		EmpID:   empID,
		LineUID: lineUID,
	}

	mockRedisCache := mock_cache.NewMockRedisCache(s.mockCtrl)
	mockEmployeeRepository := mock_dao.NewMockEmployeeRepository(s.mockCtrl)

	mockEmployeeRepository.EXPECT().
		UpsertEmployeeByIDs(employee).
		Return(exc.ErrUpsertEmployeeIDs).Times(1)

	employeeSvc := NewEmployeeService(mockEmployeeRepository, mockRedisCache, s.dummyConfig)
	err := employeeSvc.UpsertEmployeeByIDs(empID, lineUID)
	assert.EqualError(s.T(), err, exc.ErrUpsertEmployeeIDs.Error())
}
