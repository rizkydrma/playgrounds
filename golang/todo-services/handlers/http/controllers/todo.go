package controllers

import (
	"todo-services/database"
	"todo-services/handlers/http/payload/request"
	"todo-services/handlers/http/payload/response"
	"todo-services/lib"
	"todo-services/models"

	"github.com/gofiber/fiber/v2"
)

func TodoGetAll(ctx *fiber.Ctx) error {
	var todos []models.Todo

	// GET USER ID
	userId, errToken := lib.GetUserIdFromToken(ctx)
	if errToken != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
				Message: errToken.Error(),
				Code: response.UNAUTHORIZED_CODE,
		})
	}
	
	if err := database.DB.DB.Find(&todos, "user_id = ? AND deleted_at = 0", userId).Error; err != nil {
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
	var todo models.Todo
	todoId := ctx.Params("id")

	// GET USER ID
	userId, errToken := lib.GetUserIdFromToken(ctx)
	if errToken != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
				Message: errToken.Error(),
				Code: response.UNAUTHORIZED_CODE,
		})
	}

	err := database.DB.DB.First(&todo, "id = ? AND user_id = ? AND deleted_at = 0", todoId, userId).Error

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

	// GET USER ID
	userId, errToken := lib.GetUserIdFromToken(ctx)
	if errToken != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
				Message: errToken.Error(),
				Code: response.UNAUTHORIZED_CODE,
		})
	}

	 if err := lib.Validate(ctx, todo); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(
				response.BaseResponse{
					Message: err.Error(),
					Code: response.INVALID_REQUEST_PAYLOAD_CODE,
				})
	 }


	newTodo := models.Todo{
		Title: todo.Title,
		Description: todo.Description,
		UserId: userId,
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

	// GET USER ID
	userId, errToken := lib.GetUserIdFromToken(ctx)
	if errToken != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
				Message: errToken.Error(),
				Code: response.UNAUTHORIZED_CODE,
		})
	}
	
	if err := ctx.BodyParser(todoRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Message: response.INVALID_REQUEST_PAYLOAD_MESSAGE,
			Code: response.INVALID_REQUEST_PAYLOAD_CODE,
		})
	}

	if err := lib.Validate(ctx, todoRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.BaseResponse{
				Message: err.Error(),
				Code: response.INVALID_REQUEST_PAYLOAD_CODE,
			})
	}

	todoId := ctx.Params("id")

	// CHECK IF TODO EXIST
	var todo models.Todo
	if err := database.DB.DB.First(&todo, "id = ? AND user_id = ? AND deleted_at = 0", todoId, userId).Error; err != nil {
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

	// GET USER ID
	userId, errToken := lib.GetUserIdFromToken(ctx)
	if errToken != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
				Message: errToken.Error(),
				Code: response.UNAUTHORIZED_CODE,
		})
	}

	if err := ctx.BodyParser(todoRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Message: response.INVALID_REQUEST_PAYLOAD_MESSAGE,
			Code: response.INVALID_REQUEST_PAYLOAD_CODE,
		})
	}

	if err := lib.Validate(ctx,todoRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.BaseResponse{
				Message: err.Error(),
				Code: response.INVALID_REQUEST_PAYLOAD_CODE,
			})
	}

	todoId := ctx.Params("id")
	// CHECK IF TODO EXIST
	var todo models.Todo
	if err := database.DB.DB.First(&todo, "id = ? AND user_id = ? AND deleted_at = 0", todoId, userId).Error; err != nil {
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

	// GET USER ID
	userId, errToken := lib.GetUserIdFromToken(ctx)
	if errToken != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
				Message: errToken.Error(),
				Code: response.UNAUTHORIZED_CODE,
		})
	}

	if err := database.DB.DB.First(&todo, "id = ? AND user_id = ? AND deleted_at = 0", todoId, userId).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Message: response.NOT_FOUND_MESSAGE,
			Code: response.NOT_FOUND_CODE,
		})
	}

	todo.DeletedAt = 1
	if err := database.DB.DB.Save(&todo).Error; err != nil {
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