package todo

import (
	"firstapp/cmd/app"
	"firstapp/module/todo/controller"
	"firstapp/module/todo/repository"
	"firstapp/module/todo/service"
)

func Init(app app.Application) {
	repo := repository.NewTodoRepository(app.Postgres)
	serv := service.NewTodoService(repo)
	cont := controller.NewTodoController(app.Response, serv)

	r := app.Router.Group("/todos")
	r.Get("/", cont.Index)
	r.Post("/", cont.Store)
	r.Get("/:id", cont.Get)
	r.Put("/:id", cont.Update)
	r.Delete("/:id", cont.Delete)
}