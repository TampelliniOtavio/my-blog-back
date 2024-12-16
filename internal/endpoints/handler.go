package endpoints

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/contract"
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	authendpoint "github.com/TampelliniOtavio/my-blog-back/internal/endpoints/auth-endpoint"
	postendpoint "github.com/TampelliniOtavio/my-blog-back/internal/endpoints/post-endpoint"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
)

func DefineRoutes (app *fiber.App, db *sqlx.DB) {
    app.Use(logger.New())
    app.Use(recover.New(recover.Config{}))

    api := app.Group("/api")

    helper := &contract.HandlerEssentials{
        DB: db,
        Api: api,
    }

    authHandler := authendpoint.Handler{
        Service: &auth.ServiceImp{
            Repository: &database.AuthRepository{
                DB: db,
            },
        },
        Helper: helper,
    }

    authHandler.DefineRoutes()

	postHandler := postendpoint.Handler{
		Service: &post.ServiceImp{
			Repository: &database.PostRepository{
				DB: db,
			},
		},
		Helper: helper,
	}

	postHandler.DefineRoutes()
}
