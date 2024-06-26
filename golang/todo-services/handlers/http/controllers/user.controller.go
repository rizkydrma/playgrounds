package controllers

import (
	"todo-services/database"
	"todo-services/handlers/http/payload/request"
	"todo-services/handlers/http/payload/response"
	"todo-services/lib"
	"todo-services/models"
	"todo-services/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService services.UserService
}

func NewUserController () UserController {
	return UserController{}
}


func (c *UserController) GetAll(ctx *fiber.Ctx) error {

	users, err := c.userService.GetAll()
	if  err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Message: response.FAILED_GET_DATA_MESSAGE,
			Code: response.FAILED_GET_DATA_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: users,
	})
}

func (c *UserController) GetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	user, err := c.userService.GetById(userId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Message: response.NOT_FOUND_MESSAGE,
			Code: response.NOT_FOUND_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: user,
	})
}

func (c *UserController) UpdateById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	userReq := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(userReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Message: response.INVALID_REQUEST_PAYLOAD_MESSAGE,
			Code: response.INVALID_REQUEST_PAYLOAD_CODE,
		})
	}

	if err := lib.Validate(ctx, userReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.BaseResponse{
				Message: err.Error(),
				Code: response.INVALID_REQUEST_PAYLOAD_CODE,
			})
	}

	userResponse, err := c.userService.Update(*userReq, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Message: response.INTERNAL_SERVER_ERROR_MESSAGE,
			Code: response.INTERNAL_SERVER_ERROR_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: userResponse,
	})
}

func (c *UserController) UpdateEmail(ctx *fiber.Ctx) error {
	userId := ctx.Params("id");
	userReq := new(request.UserUpdateEmailRequest)

	if err := ctx.BodyParser(userReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Message: response.INVALID_REQUEST_PAYLOAD_MESSAGE,
			Code: response.INVALID_REQUEST_PAYLOAD_CODE,
		})
	}

	if err := lib.Validate(ctx, userReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.BaseResponse{
				Message: err.Error(),
				Code: response.INVALID_REQUEST_PAYLOAD_CODE,
			})
	}

	// CHECK IS EMAIL EXIST
	user, err := c.userService.CheckEmail(userReq.Email)
	if  err == nil {
		return ctx.Status(402).JSON(response.BaseResponse{
			Code: response.EMAIL_ALREADY_EXIST_CODE,
			Message: response.EMAIL_ALREADY_EXIST_MESSAGE,
		})
	}

	if err := database.DB.DB.First(&user, "id = ? AND deleted_at = 0", userId).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Message: response.NOT_FOUND_MESSAGE,
			Code: response.NOT_FOUND_CODE,
		})
	}

	user.Email = userReq.Email

	if err := database.DB.DB.Save(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Message: response.INTERNAL_SERVER_ERROR_MESSAGE,
			Code: response.INTERNAL_SERVER_ERROR_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: user,
	})
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
	var user models.User
	userId := ctx.Params("id")

	if err := database.DB.DB.First(&user, "id = ? AND deleted_at = 0", userId).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Message: response.NOT_FOUND_MESSAGE,
			Code: response.NOT_FOUND_CODE,
		})
	}

	user.DeletedAt = 1
	if err := database.DB.DB.Save(&user).Error; err != nil {
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