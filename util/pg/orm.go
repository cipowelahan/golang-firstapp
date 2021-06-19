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

func (uo utilOrm) Find(condition string, params ...interface{}) error {
	return uo.orm.Where(condition, params...).Select()
}

func (uo utilOrm) FindPk(id int) error {
	return uo.orm.Where("id=?", id).Select()
}

func (uo utilOrm) Paginate(limit int, page int) (resultLimit int, resultPage int, resultTotal int, err error) {
	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	err = uo.orm.
		Limit(limit).
		Offset((page - 1) * limit).
		Select()

	total := 0
	_, _ = uo.orm.QueryOne(pg.Scan(&total), "SELECT count(*) FROM ?TableName")

	resultLimit = limit
	resultPage = page
	resultTotal = total
	return
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
