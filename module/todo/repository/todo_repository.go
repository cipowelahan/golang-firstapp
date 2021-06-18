package repository

import (
	"firstapp/module/todo/model"
	"firstapp/util/pg"
)

type TodoRepository interface {
	Fetch(urlQuery *pg.UrlQuery) *pg.Paginate
	Find(id int) *model.Todo
	Store(model *model.Todo) *model.Todo
	Update(model *model.Todo) *model.Todo
	Delete(model *model.Todo)
}

type todoRepository struct {
	orm pg.Util
}

func NewTodoRepository(orm pg.Util) TodoRepository {
	if err := orm.CreateTable((*model.Todo)(nil), pg.UtilCreateTableOption{
		IfNotExists: true,
		Temp:        false,
	}); err != nil {
		panic(err)
	}

	return todoRepository{
		orm: orm,
	}
}

func (repo todoRepository) Fetch(urlQuery *pg.UrlQuery) *pg.Paginate {
	todos := new([]model.Todo)
	paginate, err := repo.orm.Orm(todos).Paginate(urlQuery)

	if err != nil {
		panic(err)
	}

	return paginate
}

func (repo todoRepository) Find(id int) *model.Todo {
	todo := new(model.Todo)

	if err := repo.orm.Orm(todo).Find(id); err != nil {
		panic(err)
	}

	return todo
}

func (repo todoRepository) Store(model *model.Todo) *model.Todo {
	if err := repo.orm.Orm(model).Insert(); err != nil {
		panic(err)
	}

	return model
}

func (repo todoRepository) Update(model *model.Todo) *model.Todo {
	if err := repo.orm.Orm(model).Update(); err != nil {
		panic(err)
	}

	return model
}

func (repo todoRepository) Delete(model *model.Todo) {
	if err := repo.orm.Orm(model).Delete(); err != nil {
		panic(err)
	}
}
