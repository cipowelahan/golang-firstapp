package user

import (
	"firstapp/util/pg"
)

type UserRepository interface {
	Fetch(urlQuery *UserUrlQuery) *UserPaginate
	Find(id int) *User
	FindWhere(condition string, params ...interface{}) *User
	Store(model *User) *User
	Update(model *User) *User
	Delete(model *User)
}

type userRepository struct {
	orm pg.Util
}

func NewUserRepository(orm pg.Util) UserRepository {
	if err := orm.CreateTable((*User)(nil), pg.UtilCreateTableOption{
		IfNotExists: true,
		Temp:        false,
	}); err != nil {
		panic(err)
	}

	return userRepository{
		orm: orm,
	}
}

func (repo userRepository) Fetch(urlQuery *UserUrlQuery) *UserPaginate {
	users := new([]User)
	limit, page, total, err := repo.orm.Orm(users).Paginate(urlQuery.Limit, urlQuery.Page)

	if err != nil {
		panic(err)
	}

	return &UserPaginate{
		Data:  users,
		Total: total,
		Limit: limit,
		Page:  page,
	}
}

func (repo userRepository) Find(id int) *User {
	user := new(User)

	if err := repo.orm.Orm(user).FindPk(id); err != nil {
		panic(err)
	}

	return user
}

func (repo userRepository) FindWhere(condition string, params ...interface{}) *User {
	user := new(User)

	if err := repo.orm.Orm(user).Where(condition, params...).Select(); err != nil {
		panic(err)
	}

	return user
}

func (repo userRepository) Store(model *User) *User {
	if err := repo.orm.Orm(model).Insert(); err != nil {
		panic(err)
	}

	return model
}

func (repo userRepository) Update(model *User) *User {
	if err := repo.orm.Orm(model).Update(); err != nil {
		panic(err)
	}

	return model
}

func (repo userRepository) Delete(model *User) {
	if err := repo.orm.Orm(model).Delete(); err != nil {
		panic(err)
	}
}
