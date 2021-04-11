package server

import (
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/cfs-service/server/handlers"
	"gitlab.com/cfs-service/store"
)

func Start(port uint64, s store.IStore) error {
	// Initilize handlers
	handlers.Initialize(s)

	app := fiber.New()

	// app.Handler("/healthz")

	// Authorization
	app.Use(func(c *fiber.Ctx) error {
		c.Context().SetUserValue("agency_id", "4f9b99eb-490a-484e-bade-15e3841dfda9")
		return c.Next()
	})

	apiV1 := app.Group("/api/v1")

	// POST event
	apiV1.Post("/events", handlers.PostEvent)

	apiV1.Get("/events", handlers.GetEvents)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))

	return nil
}
