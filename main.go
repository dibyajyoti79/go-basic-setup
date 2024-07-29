package main

import (
	"log"
	"os"

	"pripgo/config"
	"pripgo/database"
	"pripgo/middleware"
	"pripgo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load environment variables
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to MongoDB
	database.ConnectDB()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorMiddleware,
	})

	// Set up the logger middleware with a custom format
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${method} ${path} ${status} ${latency} ${ip}\n",
		Output:     os.Stdout,
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Local",
	}))

	// Setup routes
	routes.Setup(app)

	// Start server
	log.Fatal(app.Listen(":8000"))
}
