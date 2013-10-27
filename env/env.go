package env

import (
  "os"
)

func init() {
  os.Setenv("DATABASE_URL", "mongodb://localhost")
  os.Setenv("DATABASE_NAME", "confab_dev")
}

func Get(name string) string {
  return os.Getenv(name)
}
