package jwt

import (
	"firstapp/util/env"
	"time"

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
}

type util struct {
	env env.Util
}

func Init(env env.Util) Util {
	return util{
		env: env,
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
