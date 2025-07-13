package main

import (
	"todo-app/config"
	"todo-app/routes"
)

func main() {
	config.InitDB()

	route := routes.SetupRoute()

	route.Run(":8080")
}
