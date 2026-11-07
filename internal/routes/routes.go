package routes

import (
	"e-piket-be/internal/handlers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func SetupRoutes(app *fiber.App) {
	// Global Middlewares
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin, Content-Type, Accept, Authorization"},
		AllowMethods: []string{"GET, POST, PUT, DELETE, OPTIONS"},
	}))

	// Base welcome route
	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"app":     "e-piket-be",
			"status":  "running",
			"version": "1.0.0",
		})
	})

	// Handlers
	healthHandler := handlers.NewHealthHandler()
	userHandler := handlers.NewUserHandler()

	// Health check
	app.Get("/health", healthHandler.Check)

	// API v1 routes
	api := app.Group("/api/v1")
	{
		users := api.Group("/users")
		users.Get("/", userHandler.GetAll)
		users.Post("/", userHandler.Create)
	}
}
