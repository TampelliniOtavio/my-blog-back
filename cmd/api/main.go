package main

import (
	"fmt"

	"github.com/TampelliniOtavio/my-blog-back/internal/endpoints"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()

    db := database.NewDB()
    defer db.Close()

    if err != nil {
        panic(err)
    }

    app := fiber.New(fiber.Config{
        AppName: "Backend My Blog",
        JSONEncoder: json.Marshal,
	    JSONDecoder: json.Unmarshal,
        ErrorHandler: endpoints.ErrorHandler,
    })

    endpoints.DefineRoutes(app, db)

    data, _ := json.MarshalIndent(app.Stack(), "", "  ")
    fmt.Println(string(data))

    app.Listen(":3000")
}
