package main

import (
	"encoding/json"
	"github.com/hoisie/web"
)

func main() {
	InitDB()
	Run("localhost:3000")
}

func Run(host string) {
	web.Post("/conversation", createConversation)
	web.Run(host)
}

func createConversation(c *web.Context) {
	email1 := c.Request.PostFormValue("email1")
	email2 := c.Request.PostFormValue("email2")
	name1 := c.Request.PostFormValue("name1")
	name2 := c.Request.PostFormValue("name2")

	if len(email1) == 0 || len(email2) == 0 || len(name1) == 0 || len(name2) == 0 {
		c.Abort(400, "Params missing")
		return
	}

	conversation, err := CreateConversation(email1, name1, email2, name2)
	if err != nil {
		c.Abort(500, "")
		return
	}

	result, err := json.Marshal(conversation)
	if err != nil {
		c.Abort(500, "")
		return
	}

	c.ContentType("application/json")
	c.WriteString(string(result))
}
