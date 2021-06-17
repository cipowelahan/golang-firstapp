package service

import (
	"firstapp/module/todo/model"
	"firstapp/module/todo/repository"
)

type TodoService interface {
	Fetch() (*[]model.Todo, error)
	Find(id int) (*model.Todo, error)
	Store(model *model.Todo) (*model.Todo, error)
	Update(id int, model *model.Todo) (*model.Todo, error)
	Delete(id int) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return todoService{
		repo: repo,
	}
}

func (serv todoService) Fetch() (*[]model.Todo, error) {
	todos, err := serv.repo.Fetch()
	return todos, err
}

func (serv todoService) Find(id int) (*model.Todo, error) {
	todo, err := serv.repo.Find(id)
	return todo, err
}

func (serv todoService) Store(model *model.Todo) (*model.Todo, error) {
	todo, err := serv.repo.Store(model)
	return todo, err
}

func (serv todoService) Update(id int, model *model.Todo) (*model.Todo, error) {
	todo, err := serv.repo.Find(id)

	if err != nil {
		return nil, err
	}

	model.Id = todo.Id
	todo, err = serv.repo.Update(model)
	return todo, err
}

func (serv todoService) Delete(id int) error {
	todo, err := serv.repo.Find(id)

	if err != nil {
		return err
	}

	err = serv.repo.Delete(todo)
	return err
}
