package main_test

import (
	"github.com/codegangsta/confab/env"
	"github.com/codegangsta/confab/models"
	"os"
)

func init() {
	// override database name for testing purposes
	os.Setenv("DATABASE_NAME", "confab_test")

	println("dropping db:", env.Get("DATABASE_NAME"))
	err := models.DB.DropDatabase()
	if err != nil {
		panic(err)
	}
}
