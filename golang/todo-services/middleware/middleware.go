package middleware

import (
	"strings"
	"todo-services/lib/utils"

	"github.com/gofiber/fiber/v2"
)


func AuthMiddleware(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer "){
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing Token",
		})
	}

	bearerToken := strings.TrimPrefix(authHeader, "Bearer ")

	_, err := utils.VerifyToken(bearerToken); 
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}