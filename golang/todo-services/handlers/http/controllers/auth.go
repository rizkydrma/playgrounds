package controllers

import (
	"log"
	"time"
	"todo-services/database"
	"todo-services/handlers/http/payload/request"
	"todo-services/handlers/http/payload/response"
	"todo-services/lib"
	"todo-services/lib/utils"
	"todo-services/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *fiber.Ctx) error {
	var user models.User
	loginReq := new(request.LoginRequest)

	if err := ctx.BodyParser(loginReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Message: response.INVALID_REQUEST_PAYLOAD_MESSAGE,
			Code: response.INVALID_REQUEST_PAYLOAD_CODE,
		})
	}

	if err := lib.Validate(ctx,loginReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.BaseResponse{
				Message: err.Error(),
				Code: response.INVALID_REQUEST_PAYLOAD_CODE,
			})
	}

	// CHECK USER EXIST
	if err := database.DB.DB.First(&user, "email = ? AND deleted_at = 0", loginReq.Email).Error; err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
			Message: response.UNAUTHORIZED_MSG,
			Code: response.UNAUTHORIZED_CODE,
		})
	}
	
	// VALIDATION USER
	if err := utils.CheckPasswordHash(loginReq.Password, user.Password); !err {
		log.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
			Message: response.UNAUTHORIZED_MSG,
			Code: response.UNAUTHORIZED_CODE,
		})
	}

	log.Println("validation user", user)

	// GENERATE TOKEN
	claims := jwt.MapClaims{}
	claims["user_id"]	= user.ID
	claims["name"] 		= user.Name
	claims["email"] 	= user.Email
	claims["address"]	= user.Address
	claims["exp"]			= time.Now().Add(time.Minute * 30).Unix()

	token, err := utils.GenerateToken(&claims)

	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.BaseResponse{
			Message: response.UNAUTHORIZED_MSG,
			Code: response.UNAUTHORIZED_CODE,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}