package dao

import (
	exc "breakfaster/pkg/exception"
	"breakfaster/repository/model"
	"breakfaster/repository/schema"
	"errors"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// EmployeeRepository provides operations on employee model
type EmployeeRepository struct {
	db *gorm.DB
}

// GetEmpID queries employee ID by line UID
func (repo *EmployeeRepository) GetEmpID(lineUID string) (string, error) {
	var employee schema.SelectIDEmployee
	if err := repo.db.Model(&model.Employee{}).Where("line_uid = ?", lineUID).First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", exc.ErrEmployeeNotFound
		}
		log.Print(err)
		return "", exc.ErrGetEmployee
	}
	return employee.EmpID, nil
}

// GetLineUID queries line UID by employee ID
func (repo *EmployeeRepository) GetLineUID(EmpID string) (string, error) {
	var employee schema.SelectIDEmployee
	if err := repo.db.Model(&model.Employee{}).Where("emp_id = ?", EmpID).First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", exc.ErrEmployeeNotFound
		}
		log.Print(err)
		return "", exc.ErrGetEmployee
	}
	return employee.LineUID, nil
}

// UpsertEmployeeByIDs creates an entry in employee table by employee ID and line UID (setting the rest null)
// or replaces field(s) if there already exists same value on emp_id/line_id/access_card_nbr field,
// Note that if there exists orders that reference emp_id, then the update will fail
// if we do not cascade foreign key on updates
func (repo *EmployeeRepository) UpsertEmployeeByIDs(employee *model.Employee) error {
	if err := repo.db.Select("EmpID", "LineUID", "UpdatedAt", "CreatedAt").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "line_uid"}, {Name: "emp_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"line_uid", "emp_id", "updated_at", "created_at"}),
	}).Create(employee).Error; err != nil {
		log.Print(err)
		return exc.ErrUpsertEmployeeIDs
	}
	return nil
}

// NewEmployeeRepository is the factory for EmployeeRepository
func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}
