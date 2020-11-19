package dao

import (
	"breakfaster/config"
	exc "breakfaster/pkg/exception"
	"breakfaster/repository/model"
	"breakfaster/repository/schema"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// OrderRepositoryImpl provides operations on order model
type OrderRepositoryImpl struct {
	db     *gorm.DB
	logger *log.Entry
}

// CreateOrders creates an entry in orders table
// if the primary key (employee_emp_id + date) duplicates, then update food id field
func (repo *OrderRepositoryImpl) CreateOrders(orders *[]model.Order) error {
	if err := repo.db.Select("FoodID", "EmployeeEmpID", "Date", "UpdatedAt", "CreatedAt").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "primary"}},
		DoUpdates: clause.AssignmentColumns([]string{"food_id", "updated_at", "created_at"}),
	}).Create(orders).Error; err != nil {
		repo.logger.Error(err)
		return exc.ErrCreateOrder
	}
	return nil
}

// DeleteOrdersByLineUID deletes all orders of an employee within the given time interval by line UID
func (repo *OrderRepositoryImpl) DeleteOrdersByLineUID(lineUID string, start, end time.Time) error {
	subQuery := repo.db.Table("employees").Select("emp_id").Where("line_uid = ?", lineUID)
	if err := repo.db.Where("date BETWEEN ? AND ?", start, end).
		Where("employee_emp_id = (?)", subQuery).Delete(model.Order{}).Error; err != nil {
		repo.logger.Error(err)
		return exc.ErrDeleteOrder
	}
	return nil
}

// GetOrdersByLineUID retrieves orders according to the given line UID and time interval
func (repo *OrderRepositoryImpl) GetOrdersByLineUID(lineUID string, start, end time.Time) (*[]schema.SelectOrder, error) {
	var orders []schema.SelectOrder
	if err := repo.db.Model(&model.Order{}).Select("orders.date", "foods.food_name").
		Joins("left join foods on foods.id = orders.food_id").
		Joins("left join employees on employees.emp_id = orders.employee_emp_id").
		Where("date BETWEEN ? AND ?", start, end).Where("employees.line_uid = ?", lineUID).Scan(&orders).Error; err != nil {
		repo.logger.Error(err)
		return &[]schema.SelectOrder{}, exc.ErrGetOrder
	}
	return &orders, nil
}

// GetOrderByEmpID retrieves an order according to the given employee ID and date
func (repo *OrderRepositoryImpl) GetOrderByEmpID(empID string, date time.Time) (*schema.SelectOrderWithEmployeeEmpID, error) {
	var order schema.SelectOrderWithEmployeeEmpID
	if err := repo.db.Model(&model.Order{}).Select("orders.date", "orders.employee_emp_id", "orders.pick", "foods.food_name").
		Joins("left join foods on foods.id = orders.food_id").
		Where("date = ?", date).Where("orders.employee_emp_id = ?", empID).First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &schema.SelectOrderWithEmployeeEmpID{}, exc.ErrOrderNotFound
		}
		repo.logger.Error(err)
		return &schema.SelectOrderWithEmployeeEmpID{}, exc.ErrGetOrder
	}
	return &order, nil
}

// GetOrderByAccessCardNumber retrieves an order according to the given access card number and time
func (repo *OrderRepositoryImpl) GetOrderByAccessCardNumber(accessCardNumber string, date time.Time) (*schema.SelectOrderWithEmployeeEmpID, error) {
	var order schema.SelectOrderWithEmployeeEmpID
	if err := repo.db.Model(&model.Order{}).Select("orders.date", "orders.employee_emp_id", "orders.pick", "foods.food_name").
		Joins("left join foods on foods.id = orders.food_id").
		Joins("left join employees on employees.emp_id = orders.employee_emp_id").
		Where("date = ?", date).Where("employees.access_card_nbr = ?", accessCardNumber).First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &schema.SelectOrderWithEmployeeEmpID{}, exc.ErrOrderNotFound
		}
		repo.logger.Error(err)
		return &schema.SelectOrderWithEmployeeEmpID{}, exc.ErrGetOrder
	}
	return &order, nil
}

// UpdateOrderStatus updates the status of an order
func (repo *OrderRepositoryImpl) UpdateOrderStatus(empID string, date time.Time, pick bool, pickUpAt int64) error {
	if err := repo.db.Model(&model.Order{}).Where("employee_emp_id = ? AND date = ?", empID, date).
		Updates(model.Order{Pick: pick, PickUpAt: pickUpAt}).Error; err != nil {
		repo.logger.Error(err)
		return exc.ErrUpdateOrderStatus
	}
	return nil
}

// NewOrderRepository is the factory for OrderRepositoryImpl
func NewOrderRepository(db *gorm.DB, config *config.Config) OrderRepository {
	return &OrderRepositoryImpl{
		db: db,
		logger: config.Logger.ContextLogger.WithFields(log.Fields{
			"type": "dao:order",
		}),
	}
}
