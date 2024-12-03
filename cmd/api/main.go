package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/goccy/go-json"
)

func main() {
    app := fiber.New(fiber.Config{
        AppName: "Backend My Blog",
        JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
    })

    app.Use(logger.New())
    app.Use(recover.New(recover.Config{}))

    app.Get("/", func(c *fiber.Ctx) error {
        return c.Status(200).JSON(fiber.Map{"message": "Hello World!"})
    })

    app.Listen(":3000")
}
