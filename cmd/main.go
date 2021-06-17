package main

import (
	"firstapp/cmd/app"
	"firstapp/cmd/router"
)

func main() {
	app := app.Init()

	if err := router.Init(app); err != nil {
		panic(err)
	}
}
