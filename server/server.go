package server

import (
	"github.com/gofiber/fiber/v2"
	"sync"
)

// initialize telescope server type.
type TelescopeServer struct{}

// Start Http Server and Socket Server at same time.
func (t *TelescopeServer) RunEngine() {

	Wg := new(sync.WaitGroup)

	server := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	Wg.Add(2)

	go func() {
		server.Hooks().OnListen(func() error {
			println("-- Server has started on port :3000")
			return nil
		})
		Wg.Done()
	}()

	go func() {
		server.Listen(":3000")
		Wg.Done()
	}()

	Wg.Wait()
}
