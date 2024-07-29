package routes

import (
	"pripgo/utils"

	"github.com/gofiber/fiber/v2"
)

// Setup initializes all route modules
func Setup(app *fiber.App) {
	RegisterUserRoutes(app)

	// Catch-all route for undefined routes
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(
			utils.NewApiResponse(fiber.StatusNotFound, nil, "Route not found"),
		)
	})
}
