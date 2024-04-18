package router

import (
	"todo-services/config"
	"todo-services/handlers/http/controllers"
	"todo-services/middleware"

	"github.com/gofiber/fiber/v2"
)


func RouteInit() *fiber.App {
	app := fiber.New()
	
	api := app.Group("/api")

	// AUTH
	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.UserCreate)

	// USER
	api.Get("/users", middleware.AuthMiddleware ,controllers.UserGetAll)
	api.Get("/user/:id", middleware.AuthMiddleware,controllers.UserGetById)
	api.Patch("/user/:id", middleware.AuthMiddleware,controllers.UserUpdateById)
	api.Patch("/user/:id/change-email", middleware.AuthMiddleware, controllers.UserUpdateEmail)
	api.Delete("/user/:id", middleware.AuthMiddleware, controllers.UserDelete)

	// TODO
	api.Get("/todos", middleware.AuthMiddleware,controllers.TodoGetAll)
	api.Get("/todo/:id", middleware.AuthMiddleware,controllers.TodoGetById)
	api.Post("/todo", middleware.AuthMiddleware,controllers.TodoCreate)
	api.Patch("/todo/:id", middleware.AuthMiddleware,controllers.TodoUpdateById)
	api.Patch("/todo/:id/change-status", middleware.AuthMiddleware, controllers.TodoToggleStatusById)
	api.Delete("/todo/:id", middleware.AuthMiddleware, controllers.TodoDeleteById)
	

	// STATIC ASSET
	app.Static("/public", config.ProjectRootPath + "/public/asset")

	return app
}