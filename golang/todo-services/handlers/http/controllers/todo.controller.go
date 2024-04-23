package controllers

import (
	"todo-services/handlers/http/payload/request"
	"todo-services/handlers/http/payload/response"
	"todo-services/lib"
	"todo-services/services"

	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	todoService services.TodoService
}

func NewTodoController(todoService *services.TodoService) TodoController {
	return TodoController{
		todoService: *todoService,
	}
}

func (c *TodoController) GetAll(ctx *fiber.Ctx) error {
	// GET USER ID
	userId, errToken := lib.GetUserIdFromToken(ctx)
	if errToken != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
				Message: errToken.Error(),
				Code: response.UNAUTHORIZED_CODE,
		})
	}
	
	todos, err := c.todoService.GetAll(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
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

func (c *TodoController) GetById(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")

	// GET USER ID
	userId, errToken := lib.GetUserIdFromToken(ctx)
	if errToken != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
				Message: errToken.Error(),
				Code: response.UNAUTHORIZED_CODE,
		})
	}

	todo, err := c.todoService.GetById(todoId, userId)
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

func (c * TodoController) Create(ctx *fiber.Ctx) error {
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

	todoResponse, err := c.todoService.Create(*todo, userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.BaseResponse{
				Message: response.FAILED_STORE_DATA_MESSAGE,
				Code: response.FAILED_STORE_DATA_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: todoResponse,
	})
}

func (c *TodoController) UpdateById(ctx *fiber.Ctx) error {
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
	todoResponse, err := c.todoService.Update(*todoRequest, todoId, userId)
	if  err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Message: response.INTERNAL_SERVER_ERROR_MESSAGE,
			Code: response.INTERNAL_SERVER_ERROR_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: todoResponse,
	})
}

func (c *TodoController) ToggleStatusById(ctx *fiber.Ctx) error {
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
	todoResponse, err := c.todoService.ToggleStatus(*todoRequest, todoId, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Message: response.INTERNAL_SERVER_ERROR_MESSAGE,
			Code: response.INTERNAL_SERVER_ERROR_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: todoResponse,
	})
}

func (c *TodoController) DeleteById(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")

	// GET USER ID
	userId, errToken := lib.GetUserIdFromToken(ctx)
	if errToken != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
				Message: errToken.Error(),
				Code: response.UNAUTHORIZED_CODE,
		})
	}

	todoResponse, err := c.todoService.Delete(todoId, userId)

	if  err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Message: response.INTERNAL_SERVER_ERROR_MESSAGE,
			Code: response.INTERNAL_SERVER_ERROR_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code: response.SUCCESS_CODE,
		Message: "deleted",
		Data: todoResponse,
	})
}