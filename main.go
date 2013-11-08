package main

import (
	"encoding/json"
	"github.com/codegangsta/confab/jobs"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini/auth"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	InitDB()
	Run(":" + os.Getenv("PORT"))
}

func Run(host string) {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Welcome to confab!"
	})
	m.Post("/conversation", auth.Basic("kajabi", os.Getenv("APP_SECRET")), createConversation)
	m.Post("/mailgun/reply", mailgunReply)
	m.Run()
}

func createConversation(res http.ResponseWriter, req *http.Request) string {
	email1 := req.PostFormValue("email1")
	email2 := req.PostFormValue("email2")
	name1 := req.PostFormValue("name1")
	name2 := req.PostFormValue("name2")
	subject := req.PostFormValue("subject")
	body := req.PostFormValue("text")

	if len(email1) == 0 || len(email2) == 0 || len(name1) == 0 || len(name2) == 0 || len(subject) == 0 || len(body) == 0 {
		res.WriteHeader(400)
		return "Params missing"
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

	res.Header().Set("Content-Type", "application/json")
	return string(result)
}

func mailgunReply(res http.ResponseWriter, req *http.Request) string {
	email := req.PostFormValue("recipient")
	sp := strings.Split(email, "@")
	sp = strings.Split(sp[0], "+")
	token := sp[0]
	toIndex := sp[1]

	subject := req.PostFormValue("Subject")
	body := req.PostFormValue("stripped-text")
	conversation, err := GetConversation(token)
	if err != nil {
		return "Conversation not found"
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
		return "Conversation participant not found"
	}

	go jobs.SendMail(jobs.Mail{
		Index:   index,
		Token:   conversation.Token,
		Name:    name,
		To:      to,
		Subject: subject,
		Body:    body,
	})

	return "OK"
}
