package todo

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
	serv TodoService
}

func Init(app app.Application) {
	repo := NewTodoRepository(app.Postgres)
	serv := NewTodoService(repo)
	router := NewRouter(app.Response, serv)

	r := app.Router.Group("/todos")
	r.Get("/", router.Index)
	r.Post("/", router.Store)
	r.Get("/:id", router.Get)
	r.Put("/:id", router.Update)
	r.Delete("/:id", router.Delete)
}

func NewRouter(res response.Util, serv TodoService) Router {
	return router{
		res:  res,
		serv: serv,
	}
}

func (r router) Index(c *fiber.Ctx) error {
	urlQuery := new(TodoUrlQuery)

	if err := c.QueryParser(urlQuery); err != nil {
		panic(err)
	}

	todos := r.serv.Fetch(urlQuery)
	return r.res.Send(c, todos)
}

func (r router) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		panic(err)
	}

	todo := r.serv.Find(id)
	return r.res.Send(c, todo)
}

func (r router) Store(c *fiber.Ctx) error {
	body := new(TodoStore)

	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	if err := validation.Validate(*body); err != nil {
		return r.res.ErrorValidation(c, err)
	}

	todo := r.serv.Store(body)
	return r.res.Send(c, todo)
}

func (r router) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		panic(err)
	}

	body := new(TodoStore)

	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	if err := validation.Validate(*body); err != nil {
		return r.res.ErrorValidation(c, err)
	}

	todo := r.serv.Update(id, body)
	return r.res.Send(c, todo)
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
