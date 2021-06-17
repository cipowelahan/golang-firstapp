package repository

import (
	"firstapp/module/todo/model"
	"firstapp/util/pg"

	"github.com/go-pg/pg/v10/orm"
)

type TodoRepository interface {
	Fetch(urlQuery *pg.UrlQuery) (*pg.Paginate, error)
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

func (repo todoRepository) Fetch(urlQuery *pg.UrlQuery) (*pg.Paginate, error) {
	todos := new([]model.Todo)
	paginate, err := repo.pg.Orm(todos).Paginate(urlQuery)
	return paginate, err
}

func (repo todoRepository) Find(id int) (*model.Todo, error) {
	todo := new(model.Todo)
	err := repo.pg.Orm(todo).Find(id)
	return todo, err
}

func (repo todoRepository) Store(model *model.Todo) (*model.Todo, error) {
	_, err := repo.pg.Orm(model).Insert()
	return model, err
}

func (repo todoRepository) Update(model *model.Todo) (*model.Todo, error) {
	_, err := repo.pg.Orm(model).Update()
	return model, err
}

func (repo todoRepository) Delete(model *model.Todo) error {
	_, err := repo.pg.Orm(model).Delete()
	return err
}
