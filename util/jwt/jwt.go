package jwt

import (
	"firstapp/util/env"
	"firstapp/util/response"
	"strconv"
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
	env      env.Util
	response response.Util
}

func Init(env env.Util, response response.Util) Util {
	return util{
		env:      env,
		response: response,
	}
}

func (u util) Encode(payloads ...Payload) (string, error) {
	defaultTimeString := u.env.Get("JWT_TIME", "60")
	defaultTime, err := strconv.Atoi(defaultTimeString)
	if err != nil {
		return "", err
	}

	payload := Payload{
		UserID: 0,
		Exp:    time.Now().Add(time.Minute * time.Duration(defaultTime)).Unix(),
		Sub:    "Authentication",
		Iss:    u.env.Get("APP_NAME", "first-app-golang"),
	}

	if len(payloads) > 0 {
		payloadCostom := payloads[0]

		if payloadCostom.UserID != 0 {
			payload.UserID = payloadCostom.UserID
		}

		if payloadCostom.Exp != 0 {
			payload.Exp = payloadCostom.Exp
		}

		if payloadCostom.Sub != "" {
			payload.Sub = payloadCostom.Sub
		}

		if payloadCostom.Iss != "" {
			payload.Iss = payloadCostom.Iss
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": payload.UserID,
		"sub":     payload.Sub,
		"iss":     payload.Iss,
		"exp":     payload.Exp,
	})

	secret := u.env.Get("JWT_SECRET", "secret")
	return token.SignedString([]byte(secret))
}

func (u util) Decode(tokenString string) (map[string]interface{}, error) {
	secret := u.env.Get("JWT_SECRET", "secret")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if payload, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return payload, nil
	}

	return nil, err
}

func (u util) Middleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "null" {
		return u.response.Unauthorized(c)
	}

	arrAuthHeader := strings.Split(authHeader, " ")
	if len(arrAuthHeader) != 2 {
		return u.response.Unauthorized(c)
	}

	tokenString := arrAuthHeader[1]
	if _, err := u.Decode(tokenString); err != nil {
		return u.response.Unauthorized(c)
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
