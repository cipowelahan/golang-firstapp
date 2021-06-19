package todo

import (
	"time"
)

type Todo struct {
	tableName struct{} `pg:"todos"`
	Id        int64    `json:"id" pg:"id,fk"`
	TodoStore
	TodoAuthor
}

type TodoStore struct {
	Message string `json:"message" pg:"message,type:varchar" validate:"required"`
	IsDone  bool   `json:"is_done"`
}

type TodoAuthor struct {
	AuthorID  *int64     `json:"author_id" pg:"author_id"`
	EditorID  *int64     `json:"editor_id" pg:"editor_id"`
	CreatedAt *time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" pg:"updated_at"`
}

type TodoPaginate struct {
	Data  *[]Todo `json:"data"`
	Total int     `json:"total"`
	Limit int     `json:"limit"`
	Page  int     `json:"page"`
}

type TodoUrlQuery struct {
	Limit int `query:"limit"`
	Page  int `query:"page"`
}
