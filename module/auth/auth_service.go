package auth

import (
	"firstapp/module/user"
	"firstapp/util/bcrypt"
	"firstapp/util/jwt"
	"time"
)

type AuthService interface {
	Register(body *AuthRegister) *user.User
	Login(credential *AuthLogin) AuthToken
	User(id int64) *user.User
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

func (serv authService) Register(body *AuthRegister) *user.User {
	timeNow := time.Now()
	hashPassword, err := bcrypt.HashPassword(body.Password)

	if err != nil {
		panic(err)
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

	user := serv.userRepo.Store(data)
	return user
}

func (serv authService) Login(credential *AuthLogin) AuthToken {
	user := serv.userRepo.FindLogin("email=?", credential.Email)

	if user == nil || !bcrypt.CheckPasswordHash(credential.Password, *user.Password) {
		panic("Invalid Credential")
	}

	token, err := serv.jwt.Encode(jwt.Payload{
		UserID: user.Id,
	})

	if err != nil {
		panic(err)
	}

	return AuthToken{
		Type:  "Bearer",
		Token: token,
	}
}

func (serv authService) User(id int64) *user.User {
	user := serv.userRepo.Find(int(id))
	return user
}
