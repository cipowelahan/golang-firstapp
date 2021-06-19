package pg

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type utilOrm struct {
	orm   *orm.Query
	model interface{}
}

func (uo utilOrm) Where(condition string, params ...interface{}) utilOrm {
	uo.orm = uo.orm.Where(condition, params...)
	return uo
}

func (uo utilOrm) Select() error {
	return uo.orm.Select()
}

func (uo utilOrm) Find(id int) error {
	return uo.orm.Where("id=?", id).Select()
}

func (uo utilOrm) Paginate(urlQuery *UrlQuery) (*Paginate, error) {
	if urlQuery.Limit == 0 {
		urlQuery.Limit = 10
	}

	if urlQuery.Page == 0 {
		urlQuery.Page = 1
	}

	err := uo.orm.
		Limit(urlQuery.Limit).
		Offset((urlQuery.Page - 1) * urlQuery.Limit).
		Select()

	total := 0
	_, _ = uo.orm.QueryOne(pg.Scan(&total), "SELECT count(*) FROM ?TableName")
	return &Paginate{
		Data:  uo.model,
		Total: total,
		Limit: urlQuery.Limit,
		Page:  urlQuery.Page,
	}, err

}

func (uo utilOrm) Insert() error {
	_, err := uo.orm.Insert()
	return err
}

func (uo utilOrm) Update() error {
	_, err := uo.orm.WherePK().Update()
	return err
}

func (uo utilOrm) Delete() error {
	_, err := uo.orm.WherePK().Delete()
	return err
}
