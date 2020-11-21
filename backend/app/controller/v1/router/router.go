package router

import (
	"breakfaster/mybot"
	"breakfaster/pkg/ordertime"
	"breakfaster/service/core"
)

// Router is the router type
type Router struct {
	Bot      mybot.BreakFastBot
	timer    ordertime.OrderTimer
	foodSvc  core.FoodService
	orderSvc core.OrderService
	empSvc   core.EmployeeService
}

// NewRouter is a factory for router instance
func NewRouter(bot mybot.BreakFastBot, timer ordertime.OrderTimer,
	foodSvc core.FoodService, orderSvc core.OrderService, empSvc core.EmployeeService) *Router {
	return &Router{
		Bot:      bot,
		timer:    timer,
		foodSvc:  foodSvc,
		orderSvc: orderSvc,
		empSvc:   empSvc,
	}
}
