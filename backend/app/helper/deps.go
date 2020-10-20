package helper

import (
	c "breakfaster/config"
	"breakfaster/controller"
	"breakfaster/controller/v1/middleware"
	rv1 "breakfaster/controller/v1/router"
	"breakfaster/infrastructure/cache"
	"breakfaster/infrastructure/db"
	"breakfaster/mybot"
	"breakfaster/mybot/autoreply"
	"breakfaster/pkg/ordertime"
	"breakfaster/repository/dao"
	"breakfaster/service/core"

	"go.uber.org/dig"
)

// BuildContainer is a factory for dependency injection (DI) container
func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(c.NewConfig)

	container.Provide(db.ConnectDatabase)
	container.Provide(cache.NewMemCache)

	container.Provide(dao.NewFoodRepository)
	container.Provide(dao.NewOrderRepository)
	container.Provide(dao.NewEmployeeRepository)

	container.Provide(core.NewFoodService)
	container.Provide(core.NewOrderService)
	container.Provide(core.NewEmployeeService)

	container.Provide(ordertime.NewOrderTimer)

	container.Provide(mybot.NewBreakFaster)
	container.Provide(autoreply.NewAutoReplier)

	container.Provide(rv1.NewRouter)
	container.Provide(middleware.NewAuthChecker)

	container.Provide(controller.NewEngine)
	container.Provide(controller.NewServer)

	return container
}
