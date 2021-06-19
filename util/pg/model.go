package pg

type Paginate struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
	Limit int         `json:"limit"`
	Page  int         `json:"page"`
}

type UrlQuery struct {
	Limit int `query:"limit"`
	Page  int `query:"page"`
}
