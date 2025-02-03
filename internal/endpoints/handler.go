package endpoints

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/contract"
	authendpoint "github.com/TampelliniOtavio/my-blog-back/internal/endpoints/auth-endpoint"
	postendpoint "github.com/TampelliniOtavio/my-blog-back/internal/endpoints/post-endpoint"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func DefineRoutes(app *fiber.App, repo *database.Repository) {
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{}))

	api := app.Group("/api")

	helper := &contract.HandlerEssentials{
		Api: api,
	}

	authendpoint.DefineRoutes(repo.User, helper)

	postendpoint.DefineRoutes(repo.Post, helper)
}
