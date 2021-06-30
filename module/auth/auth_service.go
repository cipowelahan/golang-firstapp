package auth

import (
	"errors"
	"firstapp/module/user"
	"firstapp/util/bcrypt"
	"firstapp/util/jwt"
	"time"
)

type AuthService interface {
	Register(body *AuthRegister) (*user.User, error)
	Login(credential *AuthLogin) (*AuthToken, error)
	User(id int64) (*user.User, error)
}

type authService struct {
	jwt      jwt.Util
	userRepo user.UserRepository
}

func NewAuthService(jwt jwt.Util, userRepo user.UserRepository) AuthService {
	return authService{
		jwt:      jwt,
		userRepo: userRepo,
	}
}

func (serv authService) Register(body *AuthRegister) (*user.User, error) {
	timeNow := time.Now()
	hashPassword, err := bcrypt.HashPassword(body.Password)

	if err != nil {
		return nil, err
	}

	data := &user.User{
		UserData: user.UserData{
			Name:     &body.Name,
			Email:    &body.Email,
			Password: &hashPassword,
		},
		UserAuthor: user.UserAuthor{
			CreatedAt: &timeNow,
			UpdatedAt: &timeNow,
		},
	}

	return serv.userRepo.Store(data)
}

func (serv authService) Login(credential *AuthLogin) (*AuthToken, error) {
	user, err := serv.userRepo.FindLogin("email=?", credential.Email)

	if err != nil || !bcrypt.CheckPasswordHash(credential.Password, *user.Password) {
		return nil, errors.New("invalid credential")
	}

	token, err := serv.jwt.Encode(jwt.Payload{
		UserID: user.Id,
	})

	if err != nil {
		return nil, err
	}

	return &AuthToken{
		Type:  "Bearer",
		Token: token,
	}, nil
}

func (serv authService) User(id int64) (*user.User, error) {
	return serv.userRepo.Find(int(id))
}
