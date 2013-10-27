package main

import (
	"os"
)

func main() {

	// TODO these need to move outside the app
	os.Setenv("DATABASE_URL", "mongodb://localhost")
	os.Setenv("DATABASE_NAME", "confab_dev")

	Run("localhost:3000")
}
