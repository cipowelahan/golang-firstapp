package todo

import (
	"firstapp/cmd/app"
)

func Init(app app.Application) {
	repo := NewTodoRepository(app.Postgres)
	serv := NewTodoService(repo)
	cont := NewTodoController(app.Response, serv)

	r := app.Router.Group("/todos")
	r.Get("/", cont.Index)
	r.Post("/", cont.Store)
	r.Get("/:id", cont.Get)
	r.Put("/:id", cont.Update)
	r.Delete("/:id", cont.Delete)
}
