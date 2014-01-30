package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
  "github.com/codegangsta/confab/conversations"
)

func main() {
	Run()
}

func Run() {
	m := martini.Classic()

  // Middleware
	m.Use(render.Renderer())
	m.Use(DB())

	// m.Post("/mail/reply", MailFormHandler, func(f MailForm, l *log.Logger) (int, string) {
	// 	l.Println(f)
	// 	return 200, "OK"
	// })

	m.Post("/conversation", binding.Bind(conversations.Form{}), conversations.Create)

	m.Run()
}
