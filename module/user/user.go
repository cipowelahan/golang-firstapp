package user

import (
	"firstapp/cmd/app"
	"firstapp/util/response"
	"firstapp/util/validation"

	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Index(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type router struct {
	response   response.Util
	validation validation.Util
	service    UserService
}

func Init(app app.Application) {
	if err := app.Postgres.CreateTable((*User)(nil)); err != nil {
		panic(err)
	}

	repo := NewUserRepository(app.Postgres)
	service := NewUserService(repo)
	router := NewRouter(app.Response, app.Validation, service)

	r := app.Router.Group("/users")
	r.Get("/", router.Index)
	r.Post("/", router.Store)
	r.Get("/:id", router.Get)
	r.Put("/:id", router.Update)
	r.Delete("/:id", router.Delete)
}

func NewRouter(response response.Util, validation validation.Util, service UserService) Router {
	return router{
		response:   response,
		validation: validation,
		service:    service,
	}
}

func (r router) Index(c *fiber.Ctx) error {
	urlQuery := new(UserUrlQuery)
	if err := c.QueryParser(urlQuery); err != nil {
		return err
	}

	users, err := r.service.Fetch(urlQuery)
	if err != nil {
		return err
	}

	return r.response.Send(c, users)
}

func (r router) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	user, err := r.service.Find(id)
	if err != nil {
		return err
	}

	return r.response.Send(c, user)
}

func (r router) Store(c *fiber.Ctx) error {
	body := new(UserStore)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := r.validation.Validate(c, *body); err != nil {
		return err
	}

	user, err := r.service.Store(body)
	if err != nil {
		return err
	}

	return r.response.Send(c, user)
}

func (r router) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	body := new(UserUpdate)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := r.validation.Validate(c, *body); err != nil {
		return err
	}

	user, err := r.service.Update(id, body)
	if err != nil {
		return err
	}

	return r.response.Send(c, user)
}

func (r router) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	if err := r.service.Delete(id); err != nil {
		return err
	}

	return r.response.Send(c, nil, response.Config{
		Message: "Deleted",
	})
}
