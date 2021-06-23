package jwt

import (
	"firstapp/util/env"
	"firstapp/util/response"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Payload struct {
	Exp    int64
	UserID int64
	Sub    string
	Iss    string
}

type Util interface {
	Encode(payloads ...Payload) (string, error)
	Decode(tokenString string) (map[string]interface{}, error)
	Middleware(c *fiber.Ctx) error
	GetAuthorID(c *fiber.Ctx) (*int64, error)
}

type util struct {
	env env.Util
	res response.Util
}

func Init(env env.Util, res response.Util) Util {
	return util{
		env: env,
		res: res,
	}
}

func (u util) Encode(payloads ...Payload) (string, error) {
	var payload Payload

	if len(payloads) > 0 {
		payload = payloads[0]

		if payload.Exp == 0 {
			payload.Exp = time.Now().Add(time.Hour * 1).Unix()
		}

		if payload.Sub == "" {
			payload.Sub = "Authentication"
		}

		if payload.Iss == "" {
			payload.Iss = u.env.Get("APP_NAME", "first-app-golang")
		}
	} else {
		payload = Payload{
			UserID: 0,
			Exp:    time.Now().Add(time.Hour * 1).Unix(),
			Sub:    "Authentication",
			Iss:    u.env.Get("APP_NAME", "first-app-golang"),
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": payload.UserID,
		"sub":     payload.Sub,
		"iss":     payload.Iss,
		"exp":     payload.Exp,
	})

	secret := u.env.Get("APP_SECRET", "secret")

	return token.SignedString([]byte(secret))
}

func (u util) Decode(tokenString string) (map[string]interface{}, error) {
	secret := u.env.Get("APP_SECRET", "secret")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if payload, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return payload, nil
	} else {
		return nil, err
	}
}

func (u util) Middleware(c *fiber.Ctx) error {
	unauthorized := u.res.Error(c, nil, response.Config{
		Code:    401,
		Message: "Unauthorized",
	})

	authHeader := c.Get("Authorization")
	if authHeader == "null" {
		return unauthorized
	}

	arrAuthHeader := strings.Split(authHeader, " ")
	if len(arrAuthHeader) != 2 {
		return unauthorized
	}

	tokenString := arrAuthHeader[1]

	if _, err := u.Decode(tokenString); err != nil {
		return unauthorized
	}

	return c.Next()
}

func (u util) GetAuthorID(c *fiber.Ctx) (*int64, error) {
	var authorID *int64
	authHeader := c.Get("Authorization")
	arrAuthHeader := strings.Split(authHeader, " ")
	tokenString := arrAuthHeader[1]
	payload, err := u.Decode(tokenString)

	if err != nil {
		return nil, err
	}

	userID := payload["user_id"]
	switch id := userID.(type) {
	case int64:
		{
			authorID = &id
		}
	case float64:
		{
			val := id
			valInt := int(val)
			valInt64 := int64(valInt)
			authorID = &valInt64
		}
	}

	return authorID, nil
}
