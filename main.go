package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/confab/jobs"
	"github.com/hoisie/web"
	"log"
	"os"
	"strings"
)

func main() {
	InitDB()
	Run(":" + os.Getenv("PORT"))
}

func Run(host string) {
	secret := os.Getenv("APP_SECRET")
	web.Post(fmt.Sprintf("/%s/conversation", secret), createConversation)
	web.Post("/mailgun/reply", mailgunReply)
	web.Run(host)
}

func createConversation(c *web.Context) {
	email1 := c.Request.PostFormValue("email1")
	email2 := c.Request.PostFormValue("email2")
	name1 := c.Request.PostFormValue("name1")
	name2 := c.Request.PostFormValue("name2")
	subject := c.Request.PostFormValue("subject")
	body := c.Request.PostFormValue("text")

	if len(email1) == 0 || len(email2) == 0 || len(name1) == 0 || len(name2) == 0 || len(subject) == 0 || len(body) == 0 {
		c.Abort(400, "Params missing")
		return
	}

	conversation, err := CreateConversation(email1, name1, email2, name2)
	if err != nil {
		log.Panic(err)
	}

	result, err := json.Marshal(conversation)
	if err != nil {
		log.Panic(err)
	}

	// send email, this should probably be queued via sidekiq
	go jobs.SendMail(jobs.Mail{
		Index:   "1",
		Token:   conversation.Token,
		Name:    name1,
		To:      email2,
		Subject: subject,
		Body:    body,
	})

	c.ContentType("application/json")
	c.WriteString(string(result))
}

func mailgunReply(c *web.Context) {
	email := c.Request.PostFormValue("recipient")
	sp := strings.Split(email, "@")
	sp = strings.Split(sp[0], "+")
	token := sp[0]
	toIndex := sp[1]

	subject := c.Request.PostFormValue("Subject")
	body := c.Request.PostFormValue("stripped-text")
	conversation, err := GetConversation(token)
	if err != nil {
		c.Abort(200, "Conversation not found")
		return
	}

	var name string
	var to string
	var index string
	if toIndex == "1" {
		name = conversation.Name2
		to = conversation.Email1
		index = "2"
	} else if toIndex == "2" {
		name = conversation.Name1
		to = conversation.Email2
		index = "1"
	} else {
		c.Abort(200, "Conversation participant not found")
		return
	}

	go jobs.SendMail(jobs.Mail{
		Index:   index,
		Token:   conversation.Token,
		Name:    name,
		To:      to,
		Subject: subject,
		Body:    body,
	})

	c.WriteString("OK")
}
