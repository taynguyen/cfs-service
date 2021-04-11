package server

import (
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/cfs-service/server/handlers"
	"gitlab.com/cfs-service/server/middleware"
	"gitlab.com/cfs-service/store"
)

func Start(port uint64, s store.IStore) error {
	// Initilize handlers
	handlers.Initialize(s)

	app := fiber.New()

	app.Get("/healthz", func(c *fiber.Ctx) error {
		c.SendStatus(fiber.StatusOK)
		return nil
	})

	apiV1 := app.Group("/api/v1")
	apiV1.Use(middleware.AuthByHTTPSecureCookie)

	// POST event
	apiV1.Post("/events", handlers.PostEvent)

	apiV1.Get("/events", handlers.GetEvents)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))

	return nil
}
