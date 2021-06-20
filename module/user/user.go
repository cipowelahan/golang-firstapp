package user

import (
	"firstapp/cmd/app"
)

func Init(app app.Application) {
	repo := NewUserRepository(app.Postgres)
	serv := NewUserService(repo)
	cont := NewUserController(app.Response, serv)

	r := app.Router.Group("/users")
	r.Get("/", cont.Index)
	r.Post("/", cont.Store)
	r.Get("/:id", cont.Get)
	r.Put("/:id", cont.Update)
	r.Delete("/:id", cont.Delete)
}
