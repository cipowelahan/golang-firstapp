package controller

import (
	"firstapp/module/todo/model"
	"firstapp/module/todo/service"
	"firstapp/util/pg"
	"firstapp/util/response"

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
	serv service.TodoService
}

func NewTodoController(res response.Util, serv service.TodoService) TodoController {
	return todoController{
		res:  res,
		serv: serv,
	}
}

func (cont todoController) Index(c *fiber.Ctx) error {
	urlQuery := new(pg.UrlQuery)

	if err := c.QueryParser(urlQuery); err != nil {
		return cont.res.Send(c, err, nil)
	}

	todos, err := cont.serv.Fetch(urlQuery)
	return cont.res.Send(c, err, todos)
}

func (cont todoController) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return cont.res.Send(c, err, nil)
	}

	todo, err := cont.serv.Find(id)
	return cont.res.Send(c, err, todo)
}

func (cont todoController) Store(c *fiber.Ctx) error {
	todo := new(model.Todo)

	if err := c.BodyParser(todo); err != nil {
		return cont.res.Send(c, err, nil)
	}

	todo, err := cont.serv.Store(todo)
	return cont.res.Send(c, err, todo)
}

func (cont todoController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return cont.res.Send(c, err, nil)
	}

	todo := new(model.Todo)

	if err := c.BodyParser(todo); err != nil {
		return cont.res.Send(c, err, nil)
	}

	todo, err = cont.serv.Update(id, todo)
	return cont.res.Send(c, err, todo)
}

func (cont todoController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return cont.res.Send(c, err, nil)
	}

	err = cont.serv.Delete(id)
	return cont.res.Send(c, err, nil, response.Config{
		Message: "Deleted",
	})
}
