package auth

import (
	"firstapp/cmd/app"
	"firstapp/module/user"
	"firstapp/util/response"
	"firstapp/util/validation"

	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type router struct {
	res  response.Util
	serv AuthService
}

func Init(app app.Application) {
	userRepo := user.NewUserRepository(app.Postgres)
	serv := NewAuthService(app.JWT, userRepo)
	router := NewRouter(app.Response, serv)

	r := app.Router.Group("/auth")
	r.Post("/register", router.Register)
	r.Post("/login", router.Login)
}

func NewRouter(res response.Util, serv AuthService) Router {
	return router{
		res:  res,
		serv: serv,
	}
}

func (r router) Register(c *fiber.Ctx) error {
	body := new(AuthRegister)

	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	if err := validation.Validate(*body); err != nil {
		return r.res.ErrorValidation(c, err)
	}

	token := r.serv.Register(body)
	return r.res.Send(c, token)
}

func (r router) Login(c *fiber.Ctx) error {
	body := new(AuthLogin)

	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	if err := validation.Validate(*body); err != nil {
		return r.res.ErrorValidation(c, err)
	}

	token := r.serv.Login(body)
	return r.res.Send(c, token)
}
