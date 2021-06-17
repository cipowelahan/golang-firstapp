package repository

import (
	"firstapp/module/todo/model"
	"firstapp/util/pg"

	"github.com/go-pg/pg/v10/orm"
)

type TodoRepository interface {
	Fetch() (*[]model.Todo, error)
	Find(id int) (*model.Todo, error)
	Store(model *model.Todo) (*model.Todo, error)
	Update(model *model.Todo) (*model.Todo, error)
	Delete(model *model.Todo) error
}

type todoRepository struct {
	pg pg.Util
}

func NewTodoRepository(pg pg.Util) TodoRepository {
	if err := pg.DB().Model((*model.Todo)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
		Temp:        false,
	}); err != nil {
		panic(err)
	}

	return todoRepository{
		pg: pg,
	}
}

func (repo todoRepository) Fetch() (*[]model.Todo, error) {
	todos := new([]model.Todo)
	err := repo.pg.DB().Model(todos).Select()
	return todos, err
}

func (repo todoRepository) Find(id int) (*model.Todo, error) {
	todo := new(model.Todo)
	err := repo.pg.DB().Model(todo).Where("id=?", id).Select()
	return todo, err
}

func (repo todoRepository) Store(model *model.Todo) (*model.Todo, error) {
	_, err := repo.pg.DB().Model(model).Insert()
	return model, err
}

func (repo todoRepository) Update(model *model.Todo) (*model.Todo, error) {
	_, err := repo.pg.DB().Model(model).WherePK().Update()
	return model, err
}

func (repo todoRepository) Delete(model *model.Todo) error {
	_, err := repo.pg.DB().Model(model).Where("id=?", model.Id).Delete()
	return err
}
