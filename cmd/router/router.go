package router

import (
	"firstapp/cmd/app"
	"firstapp/module/todo"
	"firstapp/util/response"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init(app app.Application) error {
	var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
		errResponse := response.Config{
			Message: err.Error(),
			Code:    400,
		}

		if err == pg.ErrNoRows {
			errResponse.Message = "Row Not Found"
			errResponse.Code = 404
		}

		return app.Response.Error(c, nil, errResponse)
	}

	r := fiber.New(fiber.Config{
		ErrorHandler: DefaultErrorHandler,
	})

	r.Use(cors.New())
	r.Use(recover.New())

	app.Router = r
	todo.Init(app)

	return r.Listen(":" + app.Env.Get("APP_PORT"))
}
