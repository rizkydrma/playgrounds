package lib

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate[T any](ctx *fiber.Ctx, value T) error {
	var validate = validator.New()

	if err := validate.Struct(value); err != nil {
		return err
	}

	return nil
}