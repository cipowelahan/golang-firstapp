package service

import (
	"firstapp/module/todo/model"
	"firstapp/module/todo/repository"
	"firstapp/util/pg"
)

type TodoService interface {
	Fetch(urlQuery *pg.UrlQuery) *pg.Paginate
	Find(id int) *model.Todo
	Store(model *model.Todo) *model.Todo
	Update(id int, model *model.Todo) *model.Todo
	Delete(id int)
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return todoService{
		repo: repo,
	}
}

func (serv todoService) Fetch(urlQuery *pg.UrlQuery) *pg.Paginate {
	todos := serv.repo.Fetch(urlQuery)
	return todos
}

func (serv todoService) Find(id int) *model.Todo {
	todo := serv.repo.Find(id)
	return todo
}

func (serv todoService) Store(model *model.Todo) *model.Todo {
	todo := serv.repo.Store(model)
	return todo
}

func (serv todoService) Update(id int, model *model.Todo) *model.Todo {
	todo := serv.repo.Find(id)
	model.Id = todo.Id
	todo = serv.repo.Update(model)
	return todo
}

func (serv todoService) Delete(id int) {
	todo := serv.repo.Find(id)
	serv.repo.Delete(todo)
}
