package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Util struct {
}

func Init() Util {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error Load .env file")
	}

	return Util{}

}

func (e Util) Get(key string, defaultValues ...string) string {
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
