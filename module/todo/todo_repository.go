package todo

import (
	"firstapp/util/pg"
)

type TodoRepository interface {
	Fetch(urlQuery *TodoUrlQuery) (*TodoPaginate, error)
	Find(id int) (*Todo, error)
	Store(model *Todo) (*Todo, error)
	Update(model *Todo) (*Todo, error)
	Delete(model *Todo) error

	FetchByAuthor(urlQuery *TodoUrlQuery, authorID *int64) (*TodoPaginate, error)
	FindByAuthor(id int, authorID *int64) (*Todo, error)
}

type todoRepository struct {
	orm pg.Util
}

func NewTodoRepository(orm pg.Util) TodoRepository {
	return todoRepository{
		orm: orm,
	}
}

func (repo todoRepository) Fetch(urlQuery *TodoUrlQuery) (*TodoPaginate, error) {
	todos := new([]Todo)
	limit, page, total, err := repo.orm.
		Orm(todos).
		Search(urlQuery.Search, "message").
		Paginate(urlQuery.Limit, urlQuery.Page)

	if err != nil {
		return nil, err
	}

	return &TodoPaginate{
		Data:  todos,
		Total: total,
		Limit: limit,
		Page:  page,
	}, nil
}

func (repo todoRepository) Find(id int) (*Todo, error) {
	todo := new(Todo)

	if err := repo.orm.Orm(todo).FindPk(id); err != nil {
		return nil, err
	}

	return todo, nil
}

func (repo todoRepository) Store(model *Todo) (*Todo, error) {
	if err := repo.orm.Orm(model).Insert(); err != nil {
		return nil, err
	}

	return model, nil
}

func (repo todoRepository) Update(model *Todo) (*Todo, error) {
	if err := repo.orm.Orm(model).Update(); err != nil {
		return nil, err
	}

	return model, nil
}

func (repo todoRepository) Delete(model *Todo) error {
	if err := repo.orm.Orm(model).Delete(); err != nil {
		return err
	}

	return nil
}

func (repo todoRepository) FetchByAuthor(urlQuery *TodoUrlQuery, authorID *int64) (*TodoPaginate, error) {
	todos := new([]Todo)
	limit, page, total, err := repo.orm.
		Orm(todos).
		Where("author_id=?", *authorID).
		Search(urlQuery.Search, "message").
		Paginate(urlQuery.Limit, urlQuery.Page)

	if err != nil {
		return nil, err
	}

	return &TodoPaginate{
		Data:  todos,
		Total: total,
		Limit: limit,
		Page:  page,
	}, nil
}

func (repo todoRepository) FindByAuthor(id int, authorID *int64) (*Todo, error) {
	todo := new(Todo)

	if err := repo.orm.Orm(todo).Where("author_id=?", *authorID).FindPk(id); err != nil {
		return nil, err
	}

	return todo, nil
}
