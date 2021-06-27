package user

import (
	"firstapp/util/bcrypt"
	"time"
)

type UserService interface {
	Fetch(urlQuery *UserUrlQuery) (*UserPaginate, error)
	Find(id int) (*User, error)
	Store(body *UserStore) (*User, error)
	Update(id int, body *UserUpdate) (*User, error)
	Delete(id int) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return userService{
		repo: repo,
	}
}

func (serv userService) Fetch(urlQuery *UserUrlQuery) (*UserPaginate, error) {
	return serv.repo.Fetch(urlQuery)
}

func (serv userService) Find(id int) (*User, error) {
	return serv.repo.Find(id)
}

func (serv userService) Store(body *UserStore) (*User, error) {
	timeNow := time.Now()
	hashedPassword, err := serv.hashPassword(*body.Password)
	if err != nil {
		return nil, err
	}

	data := &User{
		UserData: UserData{
			Name:     body.Name,
			Email:    body.Email,
			Password: hashedPassword,
		},
		UserAuthor: UserAuthor{
			CreatedAt: &timeNow,
			UpdatedAt: &timeNow,
		},
	}

	return serv.repo.Store(data)
}

func (serv userService) Update(id int, body *UserUpdate) (*User, error) {
	userFind, err := serv.repo.Find(id)
	if err != nil {
		return nil, err
	}

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
		hashedPassword, err := serv.hashPassword(*body.Password)
		if err != nil {
			return nil, err
		}

		userData.Password = hashedPassword
	}

	data := &User{
		Id:         userFind.Id,
		UserData:   userData,
		UserAuthor: userFind.UserAuthor,
	}

	return serv.repo.Update(data)
}

func (serv userService) Delete(id int) error {
	user, err := serv.repo.Find(id)
	if err != nil {
		return err
	}

	return serv.repo.Delete(user)
}

func (serv userService) hashPassword(password string) (*string, error) {
	hashPassword, err := bcrypt.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &hashPassword, nil
}
