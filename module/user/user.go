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
	res  response.Util
	serv UserService
}

func Init(app app.Application) {
	repo := NewUserRepository(app.Postgres)
	serv := NewUserService(repo)
	router := NewRouter(app.Response, serv)

	r := app.Router.Group("/users")
	r.Get("/", router.Index)
	r.Post("/", router.Store)
	r.Get("/:id", router.Get)
	r.Put("/:id", router.Update)
	r.Delete("/:id", router.Delete)
}

func NewRouter(res response.Util, serv UserService) Router {
	return router{
		res:  res,
		serv: serv,
	}
}

func (r router) Index(c *fiber.Ctx) error {
	urlQuery := new(UserUrlQuery)

	if err := c.QueryParser(urlQuery); err != nil {
		panic(err)
	}

	users := r.serv.Fetch(urlQuery)
	return r.res.Send(c, users)
}

func (r router) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		panic(err)
	}

	user := r.serv.Find(id)
	return r.res.Send(c, user)
}

func (r router) Store(c *fiber.Ctx) error {
	body := new(UserStore)

	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	if err := validation.Validate(*body); err != nil {
		return r.res.ErrorValidation(c, err)
	}

	user := r.serv.Store(body)
	return r.res.Send(c, user)
}

func (r router) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		panic(err)
	}

	body := new(UserUpdate)

	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	if err := validation.Validate(*body); err != nil {
		return r.res.ErrorValidation(c, err)
	}

	user := r.serv.Update(id, body)
	return r.res.Send(c, user)
}

func (r router) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		panic(err)
	}

	r.serv.Delete(id)
	return r.res.Send(c, nil, response.Config{
		Message: "Deleted",
	})
}
