package lib

import (
	"fmt"
	"strings"
	"todo-services/lib/utils"

	"github.com/gofiber/fiber/v2"
)

func GetUserIdFromToken(ctx *fiber.Ctx) (uint, error){
	authHeader := ctx.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer "){
		return 0, fmt.Errorf("authorization header is required")
	}

	bearerToken := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := utils.DecodeToken(bearerToken)
	if err != nil {
		return 0, fmt.Errorf("invalid token")
	}

	userIdFloat, isOk := claims["user_id"].(float64)
	fmt.Println("isOK", isOk)
	if !isOk {
		return 0, fmt.Errorf("user not found/valid")
	}

	userId := uint(userIdFloat)

	return userId, nil
}