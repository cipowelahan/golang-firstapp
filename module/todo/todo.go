package todo

import (
	"firstapp/cmd/app"
	"firstapp/util/jwt"
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
	service    TodoService
	jwt        jwt.Util
}

func Init(app app.Application) {
	if err := app.Postgres.CreateTable((*Todo)(nil)); err != nil {
		panic(err)
	}

	repo := NewTodoRepository(app.Postgres)
	service := NewTodoService(repo)
	router := NewRouter(app.Response, app.Validation, service, app.JWT)

	r := app.Router.Group("/todos", app.JWT.Middleware)
	r.Get("/", router.Index)
	r.Post("/", router.Store)
	r.Get("/:id", router.Get)
	r.Put("/:id", router.Update)
	r.Delete("/:id", router.Delete)

}

func NewRouter(response response.Util, validation validation.Util, service TodoService, jwt jwt.Util) Router {
	return router{
		response:   response,
		validation: validation,
		service:    service,
		jwt:        jwt,
	}
}

func (r router) Index(c *fiber.Ctx) error {
	urlQuery := new(TodoUrlQuery)
	if err := c.QueryParser(urlQuery); err != nil {
		return err
	}

	authorID, err := r.jwt.GetAuthorID(c)
	if err != nil {
		return err
	}

	todos, err := r.service.FetchByAuthor(urlQuery, authorID)
	if err != nil {
		return err
	}

	return r.response.Send(c, todos)
}

func (r router) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	authorID, err := r.jwt.GetAuthorID(c)
	if err != nil {
		return err
	}

	todo, err := r.service.FindByAuthor(id, authorID)
	if err != nil {
		return err
	}

	return r.response.Send(c, todo)
}

func (r router) Store(c *fiber.Ctx) error {
	body := new(TodoStore)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := r.validation.Validate(c, *body); err != nil {
		return err
	}

	authorID, err := r.jwt.GetAuthorID(c)
	if err != nil {
		return err
	}

	body.AuthorID = authorID
	todo, err := r.service.Store(body)
	if err != nil {
		return err
	}

	return r.response.Send(c, todo)
}

func (r router) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	body := new(TodoStore)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := r.validation.Validate(c, *body); err != nil {
		return err
	}

	authorID, err := r.jwt.GetAuthorID(c)
	if err != nil {
		return err
	}

	body.AuthorID = authorID
	todo, err := r.service.UpdateByAuthor(id, body, authorID)
	if err != nil {
		return err
	}

	return r.response.Send(c, todo)
}

func (r router) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	authorID, err := r.jwt.GetAuthorID(c)
	if err != nil {
		return err
	}

	if err := r.service.DeleteByAuthor(id, authorID); err != nil {
		return err
	}

	return r.response.Send(c, nil, response.Config{
		Message: "Deleted",
	})
}
