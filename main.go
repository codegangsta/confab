package main

import (
	"github.com/codegangsta/confab/app"
	"github.com/codegangsta/confab/models"
)

func main() {
	app.Run(3000)
	models.Foo()
}
