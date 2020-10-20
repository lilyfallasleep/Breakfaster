package core

import (
	"breakfaster/infrastructure/cache"
	"breakfaster/repository/dao"
	"breakfaster/repository/model"
	ss "breakfaster/service/schema"
)

// EmployeeService provides methods for manipulating employees
type EmployeeService struct {
	repository *dao.EmployeeRepository
	cache      cache.GeneralCache
}

// GetEmployeeByLineUID service queries employee by line UID
func (svc *EmployeeService) GetEmployeeByLineUID(lineUID string) (*ss.JSONEmployee, error) {
	empID, err := svc.repository.GetEmpID(lineUID)
	if err != nil {
		return &ss.JSONEmployee{}, err
	}
	return &ss.JSONEmployee{
		EmpID:   empID,
		LineUID: lineUID,
	}, nil
}

// GetEmployeeByEmpID service queries employee by employee ID
func (svc *EmployeeService) GetEmployeeByEmpID(empID string) (*ss.JSONEmployee, error) {
	if lineUID, found := svc.cache.Get(empID); found {
		return &ss.JSONEmployee{
			EmpID:   empID,
			LineUID: lineUID.(string),
		}, nil
	}
	lineUID, err := svc.repository.GetLineUID(empID)
	if err != nil {
		return &ss.JSONEmployee{}, err
	}
	svc.cache.Set(empID, lineUID)

	return &ss.JSONEmployee{
		EmpID:   empID,
		LineUID: lineUID,
	}, nil
}

// UpsertEmployeeByIDs service upserts employee by employee ID or line UID
func (svc *EmployeeService) UpsertEmployeeByIDs(empID, lineUID string) error {
	employee := &model.Employee{
		EmpID:   empID,
		LineUID: lineUID,
	}
	if err := svc.repository.UpsertEmployeeByIDs(employee); err != nil {
		return err
	}
	svc.cache.Delete(empID)
	return nil
}

// NewEmployeeService is the factory for EmployeeService
func NewEmployeeService(repository *dao.EmployeeRepository, cache cache.GeneralCache) *EmployeeService {
	return &EmployeeService{
		repository: repository,
		cache:      cache,
	}
}
