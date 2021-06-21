package user

import (
	"firstapp/util/bcrypt"
	"time"
)

type UserService interface {
	Fetch(urlQuery *UserUrlQuery) *UserPaginate
	Find(id int) *User
	Store(body *UserStore) *User
	Update(id int, body *UserUpdate) *User
	Delete(id int)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return userService{
		repo: repo,
	}
}

func (serv userService) Fetch(urlQuery *UserUrlQuery) *UserPaginate {
	users := serv.repo.Fetch(urlQuery)
	return users
}

func (serv userService) Find(id int) *User {
	user := serv.repo.Find(id)
	return user
}

func (serv userService) Store(body *UserStore) *User {
	timeNow := time.Now()
	data := &User{
		UserData: UserData{
			Name:     body.Name,
			Email:    body.Email,
			Password: serv.hashPassword(*body.Password),
		},
		UserAuthor: UserAuthor{
			CreatedAt: &timeNow,
			UpdatedAt: &timeNow,
		},
	}
	user := serv.repo.Store(data)
	return user
}

func (serv userService) Update(id int, body *UserUpdate) *User {
	userFind := serv.repo.Find(id)
	timeNow := time.Now()
	userFind.UpdatedAt = &timeNow

	userData := UserData{}
	if body.Name != nil {
		userData.Name = body.Name
	}

	if body.Email != nil {
		userData.Email = body.Email
	}

	if body.Password != nil {
		userData.Password = serv.hashPassword(*body.Password)
	}

	data := &User{
		Id:         userFind.Id,
		UserData:   userData,
		UserAuthor: userFind.UserAuthor,
	}
	user := serv.repo.Update(data)
	return user
}

func (serv userService) Delete(id int) {
	user := serv.repo.Find(id)
	serv.repo.Delete(user)
}

func (serv userService) hashPassword(password string) *string {
	hashPassword, err := bcrypt.HashPassword(password)

	if err != nil {
		panic(err)
	}

	return &hashPassword
}
