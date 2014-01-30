package main

import (
	"labix.org/v2/mgo"
	"github.com/codegangsta/martini"
)

// DB Returns a martini.Handler
func DB() martini.Handler {
  session, err := mgo.Dial("mongodb://localhost")
  if err != nil {
    panic(err)
  }

  return func(c martini.Context) {
    s := session.Clone()
    c.Map(s.DB("confab"))
    defer s.Close()
    c.Next()
  }
}
