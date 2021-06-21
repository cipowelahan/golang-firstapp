package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Util interface {
	Get(key string, defaultValues ...string) string
}

type util struct {
}

func Init() Util {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error Load .env file")
	}

	return util{}

}

func (e util) Get(key string, defaultValues ...string) string {
	defVal := ""

	if len(defaultValues) > 0 {
		defVal = defaultValues[0]
	}

	if val := os.Getenv(key); val != "" {
		return val
	} else {
		return defVal
	}

}
