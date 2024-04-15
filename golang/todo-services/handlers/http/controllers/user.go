package controllers

import (
	"todo-services/database"
	"todo-services/handlers/http/payload/request"
	"todo-services/handlers/http/payload/response"
	"todo-services/lib"
	"todo-services/lib/utils"
	"todo-services/models"

	"github.com/gofiber/fiber/v2"
)


func UserGetAll(ctx *fiber.Ctx) error {
	var users []models.User

	if err := database.DB.DB.Find(&users, "deleted_at = 0").Error; err != nil {
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

func UserGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user models.User

	if err := database.DB.DB.First(&user, "id = ? AND deleted_at = 0", userId).Error; err != nil {
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

func UserCreate(ctx *fiber.Ctx) error {
	userReq := new(request.UserCreateRequest)

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
	var user models.User
	if err := database.DB.DB.First(&user, "email = ? AND deleted_at = 0", userReq.Email).Error; err == nil {
		return ctx.Status(402).JSON(response.BaseResponse{
			Code: response.EMAIL_ALREADY_EXIST_CODE,
			Message: response.EMAIL_ALREADY_EXIST_MESSAGE,
		})
	}

	newUser := models.User{
		Name: userReq.Name,
		Email: userReq.Email,
		Gender: userReq.Gender,
		Address: userReq.Address,
	}

	hashPassword, errHash := utils.HashingPassword(userReq.Password);
	if errHash != nil {
	 return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		 "message": "internal server error",
	 })
	}

	newUser.Password = hashPassword

	if err := database.DB.DB.Create(&newUser).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.BaseResponse{
				Message: response.FAILED_STORE_DATA_MESSAGE,
				Code: response.FAILED_STORE_DATA_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Message: response.SUCCESS_MESSAGE,
		Code: response.SUCCESS_CODE,
		Data: newUser,
	})
}

func UserUpdateById(ctx *fiber.Ctx) error {
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

	var user models.User
	if err := database.DB.DB.First(&user, "id = ? AND deleted_at = 0", userId).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Message: response.NOT_FOUND_MESSAGE,
			Code: response.NOT_FOUND_CODE,
		})
	}

	user.Name 		= userReq.Name
	user.Address 	= userReq.Address
	user.Gender		= userReq.Gender

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

func UserUpdateEmail(ctx *fiber.Ctx) error {
	var user models.User
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
	if err := database.DB.DB.First(&user, "email = ? AND deleted_at = 0", userReq.Email).Error; err == nil {
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

func UserDelete(ctx *fiber.Ctx) error {
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