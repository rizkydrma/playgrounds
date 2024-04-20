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
	authController := controllers.NewAuthController()
	api.Post("/login", authController.Login)
	api.Post("/register", authController.Register)

	// USER
	userController := controllers.NewUserController()
	api.Get("/users", middleware.AuthMiddleware ,userController.GetAll)
	api.Get("/user/:id", middleware.AuthMiddleware,userController.GetById)
	api.Patch("/user/:id", middleware.AuthMiddleware,userController.UpdateById)
	api.Patch("/user/:id/change-email", middleware.AuthMiddleware, userController.UpdateEmail)
	api.Delete("/user/:id", middleware.AuthMiddleware, userController.Delete)

	// TODO
	todoController := controllers.NewTodoController()
	api.Get("/todos", middleware.AuthMiddleware, todoController.GetAll)
	api.Get("/todo/:id", middleware.AuthMiddleware, todoController.GetById)
	api.Post("/todo", middleware.AuthMiddleware, todoController.Create)
	api.Patch("/todo/:id", middleware.AuthMiddleware, todoController.UpdateById)
	api.Patch("/todo/:id/change-status", middleware.AuthMiddleware, todoController.ToggleStatusById)
	api.Delete("/todo/:id", middleware.AuthMiddleware, todoController.DeleteById)
	

	// STATIC ASSET
	app.Static("/public", config.ProjectRootPath + "/public/asset")

	return app
}