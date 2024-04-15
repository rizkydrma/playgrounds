package router

import (
	"todo-services/config"
	"todo-services/handlers/http/controllers"

	"github.com/gofiber/fiber/v2"
)


func RouteInit() *fiber.App {
	app := fiber.New()
	
	api := app.Group("/api")

	// USER
	api.Get("/users", controllers.UserGetAll)
	api.Get("/user/:id", controllers.UserGetById)
	api.Post("/user/register", controllers.UserCreate)
	api.Patch("/user/:id", controllers.UserUpdateById)
	api.Patch("/user/:id/change-email", controllers.UserUpdateEmail)
	api.Delete("/user/:id", controllers.UserDelete)

	// TODO
	api.Get("/todos", controllers.TodoGetAll)
	api.Get("/todo/:id", controllers.TodoGetById)
	api.Post("/todo", controllers.TodoCreate)
	api.Patch("/todo/:id", controllers.TodoUpdateById)
	api.Patch("/todo/:id/change-status", controllers.TodoToggleStatusById)
	api.Delete("/todo/:id", controllers.TodoDeleteById)
	

	// STATIC ASSET
	app.Static("/public", config.ProjectRootPath + "/public/asset")

	return app
}