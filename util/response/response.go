package response

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Message string
	Code    int
}

type Util interface {
	Send(c *fiber.Ctx, data interface{}, configs ...Config) error
	Error(c *fiber.Ctx, data interface{}, configs ...Config) error
}

type utilResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type util struct {
	response utilResponse
}

func Init() Util {
	return util{
		response: utilResponse{},
	}
}

func (u util) getConfig(configs ...Config) Config {
	var config Config
	if len(configs) > 0 {
		config = configs[0]

		if config.Message == "" {
			config.Message = "OK"
		}

		if config.Code == 0 {
			config.Code = 200
		}
	} else {
		config = Config{
			Message: "OK",
			Code:    200,
		}
	}

	return config
}

func (u util) getConfigError(configs ...Config) Config {
	var config Config
	if len(configs) > 0 {
		config = configs[0]

		if config.Message == "" {
			config.Message = "Error Request"
		}

		if config.Code == 0 {
			config.Code = 400
		}
	} else {
		config = Config{
			Message: "Error Request",
			Code:    400,
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

func (u util) Error(c *fiber.Ctx, data interface{}, configs ...Config) error {
	config := u.getConfigError(configs...)
	u.response.Code = config.Code
	u.response.Message = config.Message
	u.response.Data = data
	return c.Status(u.response.Code).JSON(u.response)
}
