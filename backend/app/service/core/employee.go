package core

import (
	"breakfaster/config"
	"breakfaster/infrastructure/cache"
	"breakfaster/repository/dao"
	"breakfaster/repository/model"
	ss "breakfaster/service/schema"

	log "github.com/sirupsen/logrus"
)

// EmployeeService provides methods for manipulating employees
type EmployeeService struct {
	repository *dao.EmployeeRepository
	cache      *cache.RedisCache
	logger     *log.Entry
}

// GetEmployeeByLineUID service queries employee by line UID
func (svc *EmployeeService) GetEmployeeByLineUID(lineUID string) (*ss.JSONEmployee, error) {
	var empID string
	found, err := svc.cache.Get(lineUID, &empID)
	if err != nil {
		svc.logger.Error(err)
	} else if found {
		return &ss.JSONEmployee{
			EmpID:   empID,
			LineUID: lineUID,
		}, nil
	}

	empID, err = svc.repository.GetEmpID(lineUID)
	if err != nil {
		return &ss.JSONEmployee{}, err
	}
	if err := svc.cache.Set(lineUID, empID); err != nil {
		svc.logger.Error(err)
	}

	return &ss.JSONEmployee{
		EmpID:   empID,
		LineUID: lineUID,
	}, nil
}

// GetEmployeeByEmpID service queries employee by employee ID
func (svc *EmployeeService) GetEmployeeByEmpID(empID string) (*ss.JSONEmployee, error) {
	var lineUID string
	found, err := svc.cache.Get(empID, &lineUID)
	if err != nil {
		svc.logger.Error(err)
	} else if found {
		return &ss.JSONEmployee{
			EmpID:   empID,
			LineUID: lineUID,
		}, nil
	}

	lineUID, err = svc.repository.GetLineUID(empID)
	if err != nil {
		return &ss.JSONEmployee{}, err
	}
	if err := svc.cache.Set(empID, lineUID); err != nil {
		svc.logger.Error(err)
	}

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
	if err := svc.cache.ExecPipeLine(&cmds); err != nil {
		svc.logger.Error(err)
	}
	return nil
}

// NewEmployeeService is the factory for EmployeeService
func NewEmployeeService(repository *dao.EmployeeRepository, cache *cache.RedisCache, config *config.Config) *EmployeeService {
	return &EmployeeService{
		repository: repository,
		cache:      cache,
		logger: config.Logger.ContextLogger.WithFields(log.Fields{
			"type": "svc:employee",
		}),
	}
}
