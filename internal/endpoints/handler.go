package endpoints

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/contract"
	"github.com/TampelliniOtavio/my-blog-back/internal/endpoints/auth"
	"github.com/TampelliniOtavio/my-blog-back/internal/endpoints/post"
	"github.com/TampelliniOtavio/my-blog-back/internal/endpoints/user"
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

	auth.DefineRoutes(repo.User, helper)

	post.DefineRoutes(repo.Post, helper)

	user.DefineRoutes(repo.User, helper)
}
