package server

import (
	"github.com/gofiber/fiber/v2"
)

// initialize telescope server type.
type TelescopeServer struct{}

// Start Http Server and Socket Server at same time.
func (t *TelescopeServer) RunEngine() {
	server := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	server.Listen(":8080")
}
