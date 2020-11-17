package main

import (
	"breakfaster/helper"
	"breakfaster/messaging"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'send' subcommand")
		os.Exit(1)
	}
	container := helper.BuildContainer()

	switch os.Args[1] {
	case "send":
		err := container.Invoke(func(message messaging.MessageController) {
			if err := message.BroadCastMenu(); err != nil {
				fmt.Println(err)
			}
		})
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("expected 'send' subcommand")
		os.Exit(1)
	}
}
