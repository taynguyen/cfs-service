package server

import (
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/k0kubun/pp"
	"gitlab.com/cfs-service/server/handlers"
	"gitlab.com/cfs-service/store"
)

func Start(port uint64, s store.IStore) error {
	// Initilize handlers
	handlers.Initialize(s)

	app := fiber.New()

	// Authorization
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("🥇 First handler")
		c.Context().SetUserValue("my-name", "Tay Nguyen")
		return c.Next()
	})

	apiV1 := app.Group("/api/v1")

	// GET /john
	apiV1.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		val := c.Context().UserValue("my-name")
		pp.Println("val:", val)
		return c.SendString(msg) // => Hello john 👋!
	})

	// POST event
	apiV1.Post("/events", handlers.PostEvent)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))

	return nil
}
