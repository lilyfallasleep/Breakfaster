package core

import (
	"breakfaster/config"
	"breakfaster/mybot"
	exc "breakfaster/pkg/exception"
	"breakfaster/pkg/ordertime"
	"breakfaster/repository/dao"
	"breakfaster/repository/model"
	"breakfaster/service/constant"
	ss "breakfaster/service/schema"
	"time"

	log "github.com/sirupsen/logrus"
)

// OrderServiceImpl provides methods for manipulating orders
type OrderServiceImpl struct {
	orderRepository dao.OrderRepository
	empSvc          EmployeeService
	bot             mybot.BreakFastBot
	pusher          mybot.BreakFastPushBot
	timer           ordertime.OrderTimer
	logger          *log.Entry
}

// SendOrderConfirmMessage sends an order confirmation message to the employee
func (svc *OrderServiceImpl) SendOrderConfirmMessage(empID string, start, end time.Time) error {
	employee, err := svc.empSvc.GetEmployeeByEmpID(empID)
	if err != nil {
		return exc.ErrGetlineUIDWhenConfirm
	}
	lineUID := (*employee).LineUID
	confirmCard, err := svc.bot.NewConfirmCard(lineUID, start, end)
	if err != nil {
		return exc.ErrGetOrderWhenConfirm
	}
	if err := svc.pusher.SendDirectFlex(lineUID, "確認訂單", confirmCard); err != nil {
		svc.logger.Error(err)
		return exc.ErrSendMsg
	}
	return nil
}

// CreateOrders service inserts orders from an employee
func (svc *OrderServiceImpl) CreateOrders(rawOrders *ss.AllOrders) error {
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
func (svc *OrderServiceImpl) GetOrderByEmpID(empID, rawDate string) (*ss.JSONOrder, error) {
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
	if order.FoodName == "" {
		order.FoodName = constant.DeletedFoodName
	}
	return &ss.JSONOrder{
		FoodName: order.FoodName,
		EmpID:    order.EmployeeEmpID,
		Date:     order.Date.Format(constant.DateFormat),
		Pick:     order.Pick,
	}, nil
}

// GetOrderByAccessCardNumber service retrieves a daily order by employee ID
func (svc *OrderServiceImpl) GetOrderByAccessCardNumber(accessCardNumber, rawDate string) (*ss.JSONOrder, error) {
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
	if order.FoodName == "" {
		order.FoodName = constant.DeletedFoodName
	}
	return &ss.JSONOrder{
		FoodName: order.FoodName,
		EmpID:    order.EmployeeEmpID,
		Date:     order.Date.Format(constant.DateFormat),
		Pick:     order.Pick,
	}, nil
}

// SetPick service sets the pick status of an order to true
func (svc *OrderServiceImpl) SetPick(empID string, rawDate string) error {
	date, err := time.ParseInLocation(constant.DateFormat, rawDate, time.Local)
	if err != nil {
		return exc.ErrDateFormat
	}
	if err := svc.orderRepository.UpdateOrderStatus(empID, date, true, time.Now().Unix()); err != nil {
		return err
	}
	return nil
}

// NewOrderService is the factory for OrderServiceImpl
func NewOrderService(orderRepository dao.OrderRepository, empSvc EmployeeService,
	bot mybot.BreakFastBot, pusher mybot.BreakFastPushBot, timer ordertime.OrderTimer, config *config.Config) OrderService {
	return &OrderServiceImpl{
		orderRepository: orderRepository,
		empSvc:          empSvc,
		bot:             bot,
		pusher:          pusher,
		timer:           timer,
		logger: config.Logger.ContextLogger.WithFields(log.Fields{
			"type": "svc:order",
		}),
	}
}
