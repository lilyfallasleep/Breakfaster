package core

import (
	"breakfaster/mybot"
	exc "breakfaster/pkg/exception"
	"breakfaster/pkg/ordertime"
	"breakfaster/repository/dao"
	"breakfaster/repository/model"
	"breakfaster/service/constant"
	ss "breakfaster/service/schema"
	"time"
)

// OrderService provides methods for manipulating orders
type OrderService struct {
	orderRepository *dao.OrderRepository
	empRepository   *dao.EmployeeRepository
	bot             *mybot.BreakFaster
	timer           *ordertime.OrderTimer
}

// SendOrderConfirmMessage sends an order confirmation message to the employee
func (svc *OrderService) SendOrderConfirmMessage(empID string, start, end time.Time) error {
	lineUID, err := svc.empRepository.GetLineUID(empID)
	if err != nil {
		return exc.ErrGetlineUIDWhenConfirm
	}
	confirmCard, err := svc.bot.NewConfirmCard(lineUID, start, end)
	if err != nil {
		return exc.ErrGetOrderWhenConfirm
	}
	if err := svc.bot.SendDirectFlex(lineUID, "確認訂單", confirmCard); err != nil {
		return exc.ErrSendMsg
	}
	return nil
}

// CreateOrders service inserts orders from an employee
func (svc *OrderService) CreateOrders(rawOrders *ss.AllOrders) error {
	var orders []model.Order
	for _, rawOrder := range rawOrders.Foods {
		supplyDatetime, err := time.ParseInLocation(constant.DateFormat, rawOrder.Date, time.Local)
		if err != nil {
			return exc.ErrDateFormat
		}
		orders = append(orders, model.Order{
			FoodID:        rawOrder.FoodID,
			EmployeeEmpID: rawOrders.EmpID,
			Date:          supplyDatetime,
		})
	}
	if err := svc.orderRepository.CreateOrders(&orders); err != nil {
		return err
	}

	start, end := svc.timer.GetNextWeekInterval()
	if err := svc.SendOrderConfirmMessage(rawOrders.EmpID, start, end); err != nil {
		return err
	}

	return nil
}

// GetOrderByEmpID service retrieves a daily order by employee ID
func (svc *OrderService) GetOrderByEmpID(empID, rawDate string) (*ss.JSONOrder, error) {
	var date time.Time
	var err error
	date, err = time.ParseInLocation(constant.DateFormat, rawDate, time.Local)
	if err != nil {
		return &ss.JSONOrder{}, exc.ErrDateFormat
	}

	order, err := svc.orderRepository.GetOrderByEmpID(empID, date)
	if err != nil {
		return &ss.JSONOrder{}, err
	}
	return &ss.JSONOrder{
		FoodName: order.FoodName,
		EmpID:    order.EmployeeEmpID,
		Date:     order.Date.Format(constant.DateFormat),
		Pick:     order.Pick,
	}, nil
}

// GetOrderByAccessCardNumber service retrieves a daily order by employee ID
func (svc *OrderService) GetOrderByAccessCardNumber(accessCardNumber, rawDate string) (*ss.JSONOrder, error) {
	var date time.Time
	var err error
	date, err = time.ParseInLocation(constant.DateFormat, rawDate, time.Local)
	if err != nil {
		return &ss.JSONOrder{}, exc.ErrDateFormat
	}

	order, err := svc.orderRepository.GetOrderByAccessCardNumber(accessCardNumber, date)
	if err != nil {
		return &ss.JSONOrder{}, err
	}
	return &ss.JSONOrder{
		FoodName: order.FoodName,
		EmpID:    order.EmployeeEmpID,
		Date:     order.Date.Format(constant.DateFormat),
		Pick:     order.Pick,
	}, nil
}

// SetPick service sets the pick status of an order to true
func (svc *OrderService) SetPick(empID string, rawDate string) error {
	date, err := time.ParseInLocation(constant.DateFormat, rawDate, time.Local)
	if err != nil {
		return exc.ErrDateFormat
	}
	if err := svc.orderRepository.UpdateOrderStatus(empID, date, true); err != nil {
		return err
	}
	return nil
}

// NewOrderService is the factory for OrderService
func NewOrderService(orderRepository *dao.OrderRepository, empRepository *dao.EmployeeRepository,
	bot *mybot.BreakFaster, timer *ordertime.OrderTimer) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		empRepository:   empRepository,
		bot:             bot,
		timer:           timer,
	}
}
