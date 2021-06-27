package todo

import "time"

type TodoService interface {
	Fetch(urlQuery *TodoUrlQuery) (*TodoPaginate, error)
	Find(id int) (*Todo, error)
	Store(body *TodoStore) (*Todo, error)
	Update(id int, body *TodoStore) (*Todo, error)
	Delete(id int) error

	FetchByAuthor(urlQuery *TodoUrlQuery, authorID *int64) (*TodoPaginate, error)
	FindByAuthor(id int, authorID *int64) (*Todo, error)
	UpdateByAuthor(id int, body *TodoStore, authorID *int64) (*Todo, error)
	DeleteByAuthor(id int, authorID *int64) error
}

type todoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) TodoService {
	return todoService{
		repo: repo,
	}
}

func (serv todoService) Fetch(urlQuery *TodoUrlQuery) (*TodoPaginate, error) {
	return serv.repo.Fetch(urlQuery)
}

func (serv todoService) Find(id int) (*Todo, error) {
	return serv.repo.Find(id)
}

func (serv todoService) Store(body *TodoStore) (*Todo, error) {
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

	return serv.repo.Store(data)
}

func (serv todoService) Update(id int, body *TodoStore) (*Todo, error) {
	todoFind, err := serv.repo.Find(id)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	todoFind.EditorID = body.AuthorID
	todoFind.UpdatedAt = &timeNow
	data := &Todo{
		Id:         todoFind.Id,
		TodoData:   body.TodoData,
		TodoAuthor: todoFind.TodoAuthor,
	}

	return serv.repo.Update(data)
}

func (serv todoService) Delete(id int) error {
	todo, err := serv.repo.Find(id)
	if err != nil {
		return err
	}

	return serv.repo.Delete(todo)
}

func (serv todoService) FetchByAuthor(urlQuery *TodoUrlQuery, authorID *int64) (*TodoPaginate, error) {
	return serv.repo.FetchByAuthor(urlQuery, authorID)
}

func (serv todoService) FindByAuthor(id int, authorID *int64) (*Todo, error) {
	return serv.repo.FindByAuthor(id, authorID)
}

func (serv todoService) UpdateByAuthor(id int, body *TodoStore, authorID *int64) (*Todo, error) {
	todoFind, err := serv.repo.FindByAuthor(id, authorID)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	todoFind.EditorID = body.AuthorID
	todoFind.UpdatedAt = &timeNow
	data := &Todo{
		Id:         todoFind.Id,
		TodoData:   body.TodoData,
		TodoAuthor: todoFind.TodoAuthor,
	}

	return serv.repo.Update(data)
}

func (serv todoService) DeleteByAuthor(id int, authorID *int64) error {
	todo, err := serv.repo.FindByAuthor(id, authorID)
	if err != nil {
		return err
	}

	return serv.repo.Delete(todo)
}
