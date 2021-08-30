package main

import (
	"github.com/gofiber/fiber/v2"
)

// Handler functions

func apiHandler(ctx *fiber.Ctx) error {
	return ctx.Next()
}

func v1Handler(ctx *fiber.Ctx) error {
	ctx.Set("Version", "v1")
	return ctx.Next()
}

func helloWorldHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello World")
}

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Create `/api` group
	api := app.Group("/api", apiHandler)

	// Create `/v1` group within `/api` to get `/api/v1`
	v1 := api.Group("/v1", v1Handler)

	// A simple hello world request at `/api/v1/hello-world`
	v1.Get("/hello-world", helloWorldHandler)

	// Listen for HTTP requests at `http://localhost:3000`
	app.Listen(":3000")
}
