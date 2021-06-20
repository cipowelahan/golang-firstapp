package user

import (
	"firstapp/util/response"
	"firstapp/util/validation"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Index(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type userController struct {
	res  response.Util
	serv UserService
}

func NewUserController(res response.Util, serv UserService) UserController {
	return userController{
		res:  res,
		serv: serv,
	}
}

func (cont userController) Index(c *fiber.Ctx) error {
	urlQuery := new(UserUrlQuery)

	if err := c.QueryParser(urlQuery); err != nil {
		panic(err)
	}

	users := cont.serv.Fetch(urlQuery)
	return cont.res.Send(c, users)
}

func (cont userController) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		panic(err)
	}

	user := cont.serv.Find(id)
	return cont.res.Send(c, user)
}

func (cont userController) Store(c *fiber.Ctx) error {
	body := new(UserStore)

	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	if err := validation.Validate(*body); err != nil {
		return cont.res.ErrorValidation(c, err)
	}

	user := cont.serv.Store(body)
	return cont.res.Send(c, user)
}

func (cont userController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		panic(err)
	}

	body := new(UserUpdate)

	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	if err := validation.Validate(*body); err != nil {
		return cont.res.ErrorValidation(c, err)
	}

	user := cont.serv.Update(id, body)
	return cont.res.Send(c, user)
}

func (cont userController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		panic(err)
	}

	cont.serv.Delete(id)
	return cont.res.Send(c, nil, response.Config{
		Message: "Deleted",
	})
}
