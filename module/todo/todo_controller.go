package todo

import (
	"firstapp/util/response"
	"firstapp/util/validation"

	"github.com/gofiber/fiber/v2"
)

type TodoController interface {
	Index(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type todoController struct {
	res  response.Util
	serv TodoService
}

func NewTodoController(res response.Util, serv TodoService) TodoController {
	return todoController{
		res:  res,
		serv: serv,
	}
}

func (cont todoController) Index(c *fiber.Ctx) error {
	urlQuery := new(TodoUrlQuery)

	if err := c.QueryParser(urlQuery); err != nil {
		panic(err)
	}

	todos := cont.serv.Fetch(urlQuery)
	return cont.res.Send(c, todos)
}

func (cont todoController) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		panic(err)
	}

	todo := cont.serv.Find(id)
	return cont.res.Send(c, todo)
}

func (cont todoController) Store(c *fiber.Ctx) error {
	body := new(TodoStore)

	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	if err := validation.Validate(*body); err != nil {
		return cont.res.ErrorValidation(c, err)
	}

	todo := cont.serv.Store(body)
	return cont.res.Send(c, todo)
}

func (cont todoController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		panic(err)
	}

	body := new(TodoStore)

	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	if err := validation.Validate(*body); err != nil {
		return cont.res.ErrorValidation(c, err)
	}

	todo := cont.serv.Update(id, body)
	return cont.res.Send(c, todo)
}

func (cont todoController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		panic(err)
	}

	cont.serv.Delete(id)
	return cont.res.Send(c, nil, response.Config{
		Message: "Deleted",
	})
}
