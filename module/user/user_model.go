package user

import (
	"time"
)

type User struct {
	tableName struct{} `pg:"users"`
	Id        int64    `json:"id" pg:"id,fk"`
	UserData
	UserAuthor
}

type UserData struct {
	Name     *string `json:"name" pg:"name,type:varchar"`
	Email    *string `json:"email" pg:"email,type:varchar"`
	Password *string `json:"-" pg:"password,type:varchar"`
}

type UserStore struct {
	Name     *string `validate:"required"`
	Email    *string `validate:"required,email"`
	Password *string `validate:"required,min=6,max=12"`
}

type UserUpdate struct {
	Name     *string `validate:"omitempty"`
	Email    *string `validate:"omitempty,email"`
	Password *string `validate:"omitempty,min=6,max=12"`
}

type UserAuthor struct {
	AuthorID  *int64     `json:"author_id" pg:"author_id"`
	EditorID  *int64     `json:"editor_id" pg:"editor_id"`
	CreatedAt *time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" pg:"updated_at"`
}

type UserPaginate struct {
	Data  *[]User `json:"data"`
	Total int     `json:"total"`
	Limit int     `json:"limit"`
	Page  int     `json:"page"`
}

type UserUrlQuery struct {
	Limit int `query:"limit"`
	Page  int `query:"page"`
}
