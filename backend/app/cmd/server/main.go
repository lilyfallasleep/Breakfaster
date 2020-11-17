package main

import (
	"breakfaster/controller"
	_ "breakfaster/docs"
	"breakfaster/helper"
	"fmt"
	"os"

	"gorm.io/gorm"

	"breakfaster/repository/model"
	"breakfaster/service/constant"
)

// @title Breakfaster
// @version 1.0.0
// @description Breakfast Ordering System @LINE

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'serve', 'migrate', or 'dropallorder' subcommands")
		os.Exit(1)
	}
	container := helper.BuildContainer()

	switch os.Args[1] {
	case "serve":
		err := container.Invoke(func(server *controller.Server) {
			server.Run()
		})
		if err != nil {
			fmt.Println(err)
		}
	case "migrate":
		err := container.Invoke(func(db *gorm.DB) {
			db.AutoMigrate(&model.Food{}, &model.Employee{}, &model.Order{})
			var dummyFood model.Food
			if err := db.Select("FoodName", "UpdatedAt", "CreatedAt").FirstOrCreate(&dummyFood, &model.Food{
				ID:       1,
				FoodName: constant.DummyFoodName,
			}).Error; err != nil {
				fmt.Println(err)
			}
		})
		if err != nil {
			fmt.Println(err)
		}
	case "dropallorder":
		err := container.Invoke(func(db *gorm.DB) {
			db.Migrator().DropTable(&model.Order{})
			db.AutoMigrate(&model.Order{})
		})
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("expected 'serve', 'send', 'migrate', or 'dropallorder' subcommands")
		os.Exit(1)
	}
}
