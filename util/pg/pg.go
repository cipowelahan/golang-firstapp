package pg

import (
	"firstapp/util/env"

	"github.com/go-pg/pg/v10"
)

type Util interface {
	DB() *pg.DB
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
