package app

import (
	"firstapp/util/env"
	"firstapp/util/pg"
	"firstapp/util/response"

	"github.com/gofiber/fiber/v2"
)

type Application struct {
	Env      env.Util
	Postgres pg.Util
	Response response.Util
	Router   fiber.Router
}

func Init() Application {
	env := env.Init()
	postgres := pg.Init(env)
	response := response.NewResponse()

	return Application{
		Env:      env,
		Postgres: postgres,
		Response: response,
	}
}
