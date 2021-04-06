package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html"
)

func main() {
    // Create a new engine
    engine := html.New("/app/views", ".html")

    // Pass the engine to the Views
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    app.Static("/", "/app/public")

    app.Get("/", func(c *fiber.Ctx) error {
        // Render index within layouts/main
        return c.Render("index", fiber.Map{
            "Title": "Imersão Full Cycle",
        }, "layouts/main")
    })

    log.Fatal(app.Listen(":8000"))
}
