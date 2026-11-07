package handlers

import (
	"e-piket-be/config"

	"github.com/gofiber/fiber/v3"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(c fiber.Ctx) error {
	dbStatus := "connected"
	if config.DB != nil {
		sqlDB, err := config.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			dbStatus = "disconnected"
		}
	} else {
		dbStatus = "not initialized"
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "Service is healthy",
		"database": fiber.Map{
			"status": dbStatus,
		},
	})
}
