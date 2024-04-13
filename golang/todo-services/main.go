package main

import (
	"log"
	"todo-services/config"
	"todo-services/database"
	"todo-services/handlers/router"
)

func main(){
	config := config.LoadConfig()
	
	database.DatabaseInit(config)

	app := router.RouteInit()

	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}