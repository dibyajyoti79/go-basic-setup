package middleware

import (
	"log"
	"os"

	"pripgo/utils"

	"github.com/gofiber/fiber/v2"
)

func ErrorMiddleware(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var apiErr *utils.ApiError

	if e, ok := err.(*utils.ApiError); ok {
		apiErr = e
		code = apiErr.StatusCode
	} else {
		apiErr = utils.NewApiError(code, "Internal Server Error", []string{err.Error()})
	}

	if code == fiber.StatusInternalServerError {
		log.Printf("Error: %v", err)
	}
	// Check if running in development mode
	isDevelopment := os.Getenv("ENVIRONMENT") == "development"

	response := utils.ApiResponse{
		StatusCode: apiErr.StatusCode,
		Success:    false,
		Message:    apiErr.Message,
		Data:       nil,
	}

	// Include errors and stack trace only in development mode
	if isDevelopment && code == fiber.StatusInternalServerError {
		response.Errors = apiErr.Errors
		response.Stack = apiErr.Stack
	}

	return ctx.Status(apiErr.StatusCode).JSON(response)
}
