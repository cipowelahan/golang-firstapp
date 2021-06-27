package router

import (
	"firstapp/cmd/app"
	"firstapp/module/auth"
	"firstapp/module/todo"
	"firstapp/module/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init(app app.Application) error {
	var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
		return app.Response.Error(c, err)
	}

	r := fiber.New(fiber.Config{
		ErrorHandler: DefaultErrorHandler,
	})

	r.Use(cors.New())
	r.Use(recover.New())

	app.Router = r
	user.Init(app)
	auth.Init(app)
	todo.Init(app)

	r.Use(func(c *fiber.Ctx) error {
		return app.Response.RouteNotFound(c)
	})

	return r.Listen(":" + app.Env.Get("APP_PORT"))
}
