package todo

import "time"

type TodoService interface {
	Fetch(urlQuery *TodoUrlQuery) *TodoPaginate
	Find(id int) *Todo
	Store(body *TodoStore) *Todo
	Update(id int, body *TodoStore) *Todo
	Delete(id int)
}

type todoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) TodoService {
	return todoService{
		repo: repo,
	}
}

func (serv todoService) Fetch(urlQuery *TodoUrlQuery) *TodoPaginate {
	todos := serv.repo.Fetch(urlQuery)
	return todos
}

func (serv todoService) Find(id int) *Todo {
	todo := serv.repo.Find(id)
	return todo
}

func (serv todoService) Store(body *TodoStore) *Todo {
	timeNow := time.Now()
	data := &Todo{
		TodoData: body.TodoData,
		TodoAuthor: TodoAuthor{
			AuthorID:  body.AuthorID,
			EditorID:  body.AuthorID,
			CreatedAt: &timeNow,
			UpdatedAt: &timeNow,
		},
	}
	todo := serv.repo.Store(data)
	return todo
}

func (serv todoService) Update(id int, body *TodoStore) *Todo {
	todoFind := serv.repo.Find(id)
	timeNow := time.Now()
	todoFind.EditorID = body.AuthorID
	todoFind.UpdatedAt = &timeNow
	data := &Todo{
		Id:         todoFind.Id,
		TodoData:   body.TodoData,
		TodoAuthor: todoFind.TodoAuthor,
	}
	todo := serv.repo.Update(data)
	return todo
}

func (serv todoService) Delete(id int) {
	todo := serv.repo.Find(id)
	serv.repo.Delete(todo)
}
