package router

import (
	"todo-services/handlers/http/controllers"

	"github.com/gofiber/fiber/v2"
)


func RouteInit() *fiber.App {
	app := fiber.New()
	
	api := app.Group("/api")
	api.Get("/todos", controllers.TodoGetAll)
	api.Get("/todo/:id", controllers.TodoGetById)
	api.Post("/todo", controllers.TodoCreate)
	api.Patch("/todo/:id", controllers.TodoUpdateById)
	api.Patch("/todo/:id/change-status", controllers.TodoToggleStatusById)
	api.Delete("/todo/:id", controllers.TodoDeleteById)
	

	return app
}