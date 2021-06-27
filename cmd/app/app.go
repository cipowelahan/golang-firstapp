package app

import (
	"firstapp/util/env"
	"firstapp/util/jwt"
	"firstapp/util/pg"
	"firstapp/util/response"
	"firstapp/util/validation"

	"github.com/gofiber/fiber/v2"
)

type Application struct {
	Env        env.Util
	Postgres   pg.Util
	Response   response.Util
	Validation validation.Util
	Router     fiber.Router
	JWT        jwt.Util
}

func Init() Application {
	env := env.Init()
	postgres := pg.Init(env)
	response := response.Init()
	validation := validation.Init(response)
	jwt := jwt.Init(env, response)

	return Application{
		Env:        env,
		Postgres:   postgres,
		Response:   response,
		Validation: validation,
		JWT:        jwt,
	}
}
