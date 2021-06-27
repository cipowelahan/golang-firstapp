package user

import (
	"firstapp/util/pg"
)

type UserRepository interface {
	Fetch(urlQuery *UserUrlQuery) (*UserPaginate, error)
	Find(id int) (*User, error)
	FindLogin(condition string, params ...interface{}) (*User, error)
	Store(model *User) (*User, error)
	Update(model *User) (*User, error)
	Delete(model *User) error
}

type userRepository struct {
	orm pg.Util
}

func NewUserRepository(orm pg.Util) UserRepository {
	return userRepository{
		orm: orm,
	}
}

func (repo userRepository) Fetch(urlQuery *UserUrlQuery) (*UserPaginate, error) {
	users := new([]User)
	limit, page, total, err := repo.orm.Orm(users).Paginate(urlQuery.Limit, urlQuery.Page)

	if err != nil {
		return nil, err
	}

	return &UserPaginate{
		Data:  users,
		Total: total,
		Limit: limit,
		Page:  page,
	}, err
}

func (repo userRepository) Find(id int) (*User, error) {
	user := new(User)

	if err := repo.orm.Orm(user).FindPk(id); err != nil {
		return nil, err
	}

	return user, nil
}

func (repo userRepository) FindLogin(condition string, params ...interface{}) (*User, error) {
	user := new(User)

	if err := repo.orm.Orm(user).Where(condition, params...).SelectOne(); err != nil {
		return nil, err
	}

	return user, nil
}

func (repo userRepository) Store(model *User) (*User, error) {
	if err := repo.orm.Orm(model).Insert(); err != nil {
		return nil, err
	}

	return model, nil
}

func (repo userRepository) Update(model *User) (*User, error) {
	if err := repo.orm.Orm(model).Update(); err != nil {
		return nil, err
	}

	return model, nil
}

func (repo userRepository) Delete(model *User) error {
	if err := repo.orm.Orm(model).Delete(); err != nil {
		return err
	}

	return nil
}
