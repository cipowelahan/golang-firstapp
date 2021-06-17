package router

import (
	"firstapp/cmd/app"
	"firstapp/module/todo"

	"github.com/gofiber/fiber/v2"
)

func Init(app app.Application) error {
	r := fiber.New()

	app.Router = r
	todo.Init(app)

	return r.Listen(":" + app.Env.Get("APP_PORT"))
}
