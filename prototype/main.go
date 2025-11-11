package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	// Proxy /api/users/* → http://localhost:8081/*
	app.Use("/api/users", proxy.Balancer(proxy.Config{
		Servers: []string{"http://localhost:8081"},
		ModifyRequest: func(c *fiber.Ctx) error {
			// Remove "/api/users" prefix so backend sees "/"
			c.Request().SetRequestURI("/" + c.Params("*"))
			return nil
		},
	}))

	// Proxy /api/orders/* → http://localhost:8082/*
	app.Use("/api/orders", proxy.Balancer(proxy.Config{
		Servers: []string{"http://localhost:8082"},
	}))

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Not found",
		})
	})

	log.Println("🚀 API Gateway running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
