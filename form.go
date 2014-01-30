package main

import (
  "net/http"
  "github.com/codegangsta/martini"
)

type MailForm struct {
  To string
  From string
}

func MailFormHandler(rw http.ResponseWriter, r *http.Request, c martini.Context) {
  err := r.ParseMultipartForm(int64(1024 * 1024 * 10))
  if err != nil {
    http.Error(rw, "Cannot parse form", 400)
  }

  c.Map(MailForm{
    To: r.FormValue("to"),
    From: r.FormValue("from"),
  })
}
