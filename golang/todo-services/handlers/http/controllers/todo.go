package controllers

import (
	"todo-services/database"
	"todo-services/handlers/http/payload/request"
	"todo-services/handlers/http/payload/response"
	"todo-services/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()


func TodoGetAll(ctx *fiber.Ctx) error {
	var todos []models.Todo

	err := database.DB.DB.Find(&todos, "deleted_at = 0").Error

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Message: response.FAILED_GET_DATA_MESSAGE,
			Code: response.FAILED_GET_DATA_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: todos,
	})
}

func TodoGetById(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")

	var todo models.Todo
	err := database.DB.DB.First(&todo, "id = ? AND deleted_at = 0", todoId).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Message: response.NOT_FOUND_MESSAGE,
			Code: response.NOT_FOUND_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		response.BaseResponse{
			Message: response.SUCCESS_MESSAGE,
			Code: response.SUCCESS_CODE,
			Data: todo,
		})

}

func TodoCreate(ctx *fiber.Ctx) error {
	todo := new(request.TodoRequest)
	
	if err := ctx.BodyParser(todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Message: response.INVALID_REQUEST_PAYLOAD_MESSAGE,
			Code: response.INVALID_REQUEST_PAYLOAD_CODE,
		})
	}


	if err := validate.Struct(todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.BaseResponse{
				Message: err.Error(),
				Code: response.INVALID_REQUEST_PAYLOAD_CODE,
			})
	}

	newTodo := models.Todo{
		Title: todo.Title,
		Description: todo.Description,
	}

	err := database.DB.DB.Create(&newTodo).Error; if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.BaseResponse{
				Message: response.FAILED_STORE_DATA_MESSAGE,
				Code: response.FAILED_STORE_DATA_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: newTodo,
	})
}

func TodoUpdateById(ctx *fiber.Ctx) error {
	todoRequest := new(request.TodoUpdateRequest)
	
	if err := ctx.BodyParser(todoRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Message: response.INVALID_REQUEST_PAYLOAD_MESSAGE,
			Code: response.INVALID_REQUEST_PAYLOAD_CODE,
		})
	}

	if err := validate.Struct(todoRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.BaseResponse{
				Message: err.Error(),
				Code: response.INVALID_REQUEST_PAYLOAD_CODE,
			})
	}

	todoId := ctx.Params("id")

	// CHECK IF TODO EXIST
	var todo models.Todo
	if err := database.DB.DB.First(&todo, "id = ? AND deleted_at = 0", todoId).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Message: response.NOT_FOUND_MESSAGE,
			Code: response.NOT_FOUND_CODE,
		})
	}

	// SETUP PAYLOAD UPDATE
	todo.Title = todoRequest.Title
	todo.Description = todoRequest.Description

	if err := database.DB.DB.Save(&todo).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Message: response.INTERNAL_SERVER_ERROR_MESSAGE,
			Code: response.INTERNAL_SERVER_ERROR_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: todo,
	})
}

func TodoToggleStatusById(ctx *fiber.Ctx) error {
	todoRequest := new(request.TodoToggleRequest)

	if err := ctx.BodyParser(todoRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Message: response.INVALID_REQUEST_PAYLOAD_MESSAGE,
			Code: response.INVALID_REQUEST_PAYLOAD_CODE,
		})
	}

	if err := validate.Struct(todoRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.BaseResponse{
				Message: err.Error(),
				Code: response.INVALID_REQUEST_PAYLOAD_CODE,
			})
	}

	todoId := ctx.Params("id")
	// CHECK IF TODO EXIST
	var todo models.Todo
	if err := database.DB.DB.First(&todo, "id = ? AND deleted_at = 0", todoId).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Message: response.NOT_FOUND_MESSAGE,
			Code: response.NOT_FOUND_CODE,
		})
	}

	todo.Status = todoRequest.Status
	if err := database.DB.DB.Save(&todo).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Message: response.INTERNAL_SERVER_ERROR_MESSAGE,
			Code: response.INTERNAL_SERVER_ERROR_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: todo,
	})
}

func TodoDeleteById(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")
	var todo models.Todo

	if err := database.DB.DB.First(&todo, "id = ? AND deleted_at = 0", todoId).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Message: response.NOT_FOUND_MESSAGE,
			Code: response.NOT_FOUND_CODE,
		})
	}

	if err := database.DB.DB.Delete(&todo).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Message: response.INTERNAL_SERVER_ERROR_MESSAGE,
			Code: response.INTERNAL_SERVER_ERROR_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code: response.SUCCESS_CODE,
		Message: "deleted",
	})
}