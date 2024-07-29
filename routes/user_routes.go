package routes

import (
	"pripgo/controllers"

	"github.com/gofiber/fiber/v2"
)

// RegisterUserRoutes registers routes related to user operations
func RegisterUserRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/user/:id", controllers.GetUser)
}
