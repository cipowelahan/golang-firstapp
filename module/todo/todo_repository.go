package todo

import (
	"firstapp/util/pg"
)

type TodoRepository interface {
	Fetch(urlQuery *TodoUrlQuery) *TodoPaginate
	Find(id int) *Todo
	Store(model *Todo) *Todo
	Update(model *Todo) *Todo
	Delete(model *Todo)
}

type todoRepository struct {
	orm pg.Util
}

func NewTodoRepository(orm pg.Util) TodoRepository {
	if err := orm.CreateTable((*Todo)(nil), pg.UtilCreateTableOption{
		IfNotExists: true,
		Temp:        false,
	}); err != nil {
		panic(err)
	}

	return todoRepository{
		orm: orm,
	}
}

func (repo todoRepository) Fetch(urlQuery *TodoUrlQuery) *TodoPaginate {
	todos := new([]Todo)
	limit, page, total, err := repo.orm.Orm(todos).Paginate(urlQuery.Limit, urlQuery.Page)

	if err != nil {
		panic(err)
	}

	return &TodoPaginate{
		Data:  todos,
		Total: total,
		Limit: limit,
		Page:  page,
	}
}

func (repo todoRepository) Find(id int) *Todo {
	todo := new(Todo)

	if err := repo.orm.Orm(todo).FindPk(id); err != nil {
		panic(err)
	}

	return todo
}

func (repo todoRepository) Store(model *Todo) *Todo {
	if err := repo.orm.Orm(model).Insert(); err != nil {
		panic(err)
	}

	return model
}

func (repo todoRepository) Update(model *Todo) *Todo {
	if err := repo.orm.Orm(model).Update(); err != nil {
		panic(err)
	}

	return model
}

func (repo todoRepository) Delete(model *Todo) {
	if err := repo.orm.Orm(model).Delete(); err != nil {
		panic(err)
	}
}
