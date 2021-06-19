package pg

import (
	"firstapp/util/env"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Util interface {
	DB() *pg.DB
	CreateTable(model interface{}, opt UtilCreateTableOption) error
	Scan(values ...interface{}) orm.ColumnScanner
	Orm(model interface{}) utilOrm
}

type UtilCreateTableOption struct {
	Varchar     int // replaces PostgreSQL data type `text` with `varchar(n)`
	Temp        bool
	IfNotExists bool

	// FKConstraints causes CreateTable to create foreign key constraints
	// for has one relations. ON DELETE hook can be added using tag
	// `pg:"on_delete:RESTRICT"` on foreign key field. ON UPDATE hook can be added using tag
	// `pg:"on_update:CASCADE"`
	FKConstraints bool
}

type util struct {
	pg *pg.DB
}

func Init(env env.Util) Util {

	db := pg.Connect(&pg.Options{
		Addr:     env.Get("PG_HOST", "localhost") + ":" + env.Get("PG_PORT", "5432"),
		Database: env.Get("PG_DATABASE", "postgres"),
		User:     env.Get("PG_USER", "postgres"),
		Password: env.Get("PG_PASSWORD", "postgres"),
	})

	return util{
		pg: db,
	}
}

func (u util) DB() *pg.DB {
	return u.pg
}

func (u util) CreateTable(model interface{}, opt UtilCreateTableOption) error {
	createTableoption := orm.CreateTableOptions{}
	createTableoption.Varchar = opt.Varchar
	createTableoption.IfNotExists = opt.IfNotExists
	createTableoption.Temp = opt.Temp
	createTableoption.FKConstraints = opt.FKConstraints
	return u.DB().Model(model).CreateTable(&createTableoption)
}

func (u util) Scan(values ...interface{}) orm.ColumnScanner {
	return pg.Scan(values...)
}

func (u util) Orm(model interface{}) utilOrm {
	return utilOrm{
		orm:   u.DB().Model(model),
		model: model,
	}
}
