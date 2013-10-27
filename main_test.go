package main

import (
	"os"
)

func init() {
	// override database name for testing purposes
	os.Setenv("DB_NAME", "confab_test")

	InitDB()

	println("dropping db:", os.Getenv("DB_NAME"))
	err := DB.DropDatabase()
	if err != nil {
		panic(err)
	}
}
