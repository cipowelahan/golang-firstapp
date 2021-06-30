package response

import (
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Message string
	Code    int
}

type Util interface {
	Send(c *fiber.Ctx, data interface{}, configs ...Config) error
	Error(c *fiber.Ctx, err error) error
	Unauthorized(c *fiber.Ctx) error
	RouteNotFound(c *fiber.Ctx) error
}

type UtilResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type util struct {
	response UtilResponse
}

func Init() Util {
	return util{
		response: UtilResponse{},
	}
}

func (u util) getConfig(configs ...Config) Config {
	config := Config{
		Message: "ok",
		Code:    200,
	}

	if len(configs) > 0 {
		costomConfig := configs[0]

		if costomConfig.Message != "" {
			config.Message = costomConfig.Message
		}

		if costomConfig.Code != 0 {
			config.Code = costomConfig.Code
		}
	}

	return config
}

func (u util) Send(c *fiber.Ctx, data interface{}, configs ...Config) error {
	config := u.getConfig(configs...)
	u.response.Code = config.Code
	u.response.Message = config.Message
	u.response.Data = data
	return c.Status(u.response.Code).JSON(u.response)
}

func (u util) Error(c *fiber.Ctx, err error) error {
	config := Config{
		Code:    400,
		Message: err.Error(),
	}

	if err == pg.ErrNoRows {
		config.Code = 404
		config.Message = "row not found"
	}

	return u.Send(c, nil, config)
}

func (u util) Unauthorized(c *fiber.Ctx) error {
	config := Config{
		Code:    401,
		Message: "unauthorized",
	}

	return u.Send(c, nil, config)
}

func (u util) RouteNotFound(c *fiber.Ctx) error {
	config := Config{
		Code:    401,
		Message: "route not found",
	}

	return u.Send(c, nil, config)
}
