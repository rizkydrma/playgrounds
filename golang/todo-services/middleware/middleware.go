package middleware

import (
	"todo-services/lib/utils"

	"github.com/gofiber/fiber/v2"
)


func AuthMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")

	_, err := utils.VerifyToken(token); 
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}