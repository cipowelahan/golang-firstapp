package pg

import (
	"strings"

	"github.com/go-pg/pg/v10/orm"
)

type utilOrm struct {
	orm   *orm.Query
	model interface{}
}

func (uo utilOrm) Where(condition string, params ...interface{}) utilOrm {
	if len(params) == 0 {
		return uo
	}

	uo.orm = uo.orm.Where(condition, params...)
	return uo
}

func (uo utilOrm) Search(search *string, colomns ...string) utilOrm {
	if search == nil || len(colomns) == 0 {
		return uo
	}

	for i := 0; i < len(colomns); i++ {
		colomns[i] = colomns[i] + " ilike ?"
	}

	condition := strings.Join(colomns, " or ")
	param := "%" + *search + "%"
	uo.orm = uo.orm.Where(condition, param)
	return uo
}

func (uo utilOrm) Select() error {
	return uo.orm.Select()
}

func (uo utilOrm) SelectOne() error {
	return uo.orm.Limit(1).Select()
}

func (uo utilOrm) Find(condition string, params ...interface{}) error {
	if len(params) == 0 {
		return uo.orm.Select()
	}

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

	errSelect := uo.orm.
		Limit(limit).
		Offset((page - 1) * limit).
		Select()

	if errSelect != nil {
		err = errSelect
		return
	}

	total := 0
	total, errCount := uo.orm.Count()
	if errCount != nil {
		err = errCount
		return
	}

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
