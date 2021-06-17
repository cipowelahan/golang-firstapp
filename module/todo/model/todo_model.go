package model

type Todo struct {
	tableName struct{} `pg:"todos"`
	Id        int64    `json:"id" pg:"id,fk"`
	Message   string   `json:"message" pg:"message,type:varchar"`
	IsDone    bool     `json:"is_done"`
}
