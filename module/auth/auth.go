package auth

import (
	"firstapp/cmd/app"
	"firstapp/module/user"
	"firstapp/util/jwt"
	"firstapp/util/response"
	"firstapp/util/validation"

	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	User(c *fiber.Ctx) error
}

type router struct {
	response   response.Util
	validation validation.Util
	service    AuthService
	jwt        jwt.Util
}

func Init(app app.Application) {
	userRepo := user.NewUserRepository(app.Postgres)
	service := NewAuthService(app.JWT, userRepo)
	router := NewRouter(app.Response, app.Validation, service, app.JWT)

	r := app.Router.Group("/auth")
	r.Post("/register", router.Register)
	r.Post("/login", router.Login)

	authR := r.Group("", app.JWT.Middleware)
	authR.Get("/user", router.User)
}

func NewRouter(response response.Util, validation validation.Util, service AuthService, jwt jwt.Util) Router {
	return router{
		response:   response,
		validation: validation,
		service:    service,
		jwt:        jwt,
	}
}

func (r router) Register(c *fiber.Ctx) error {
	body := new(AuthRegister)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := r.validation.Validate(c, *body); err != nil {
		return err
	}

	token, err := r.service.Register(body)
	if err != nil {
		return err
	}

	return r.response.Send(c, token)
}

func (r router) Login(c *fiber.Ctx) error {
	body := new(AuthLogin)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := r.validation.Validate(c, *body); err != nil {
		return err
	}

	token, err := r.service.Login(body)
	if err != nil {
		return err
	}

	return r.response.Send(c, token)
}

func (r router) User(c *fiber.Ctx) error {
	authorID, err := r.jwt.GetAuthorID(c)
	if err != nil {
		return err
	}

	user, err := r.service.User(*authorID)
	if err != nil {
		return err
	}

	return r.response.Send(c, user)
}
