package main

import (
	"fmt"

	"github.com/TampelliniOtavio/my-blog-back/internal/endpoints"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database"
	"github.com/fatih/color"
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
		AppName:      "Backend My Blog",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: endpoints.ErrorHandler,
	})

	endpoints.DefineRoutes(app, db)

	printRoutes(app)

	app.Listen(":3000")
}

func printRoutes(app *fiber.App) {
	index := 0

	for _, routes := range app.Stack() {
		for _, route := range routes {
			fmt.Printf(
				"%d %s %s %s\n",
				index,
				colorizeMethod(route.Method),
				route.Path,
				route.Name,
			)
			index++
		}
	}
}

func colorizeMethod(method string) string {
	green := color.New(color.Bold, color.FgGreen).SprintFunc()
	blue := color.New(color.Bold, color.FgBlue).SprintFunc()
	red := color.New(color.Bold, color.FgRed).SprintFunc()
	yellow := color.New(color.Bold, color.FgYellow).SprintFunc()
	cyan := color.New(color.Bold, color.FgCyan).SprintFunc()

	if method == "GET" || method == "HEAD" {
		return yellow(method)
	}

	if method == "POST" {
		return green(method)
	}

	if method == "PUT" || method == "PATCH" {
		return blue(method)
	}

	if method == "DELETE" {
		return red(method)
	}

	return cyan(method)
}
