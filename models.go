package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
)

var DB *mgo.Database

func InitDB() {
	session, err := mgo.Dial(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	name := os.Getenv("DB_NAME")
	println("connecting to db:", name)
	DB = session.DB(name)
	Conversations = DB.C("conversations")
}

var Conversations *mgo.Collection

type Conversation struct {
	Token  string `json:"token"`
	Email1 string `json:"email1"`
	Email2 string `json:"email2"`
}

func CreateConversation(email1 string, email2 string) (*Conversation, error) {
	token := NewUUID().String()

	c := Conversation{
		Token:  token,
		Email1: email1,
		Email2: email2,
	}

	err := Conversations.Insert(&c)

	return &c, err
}

func GetConversation(token string) (*Conversation, error) {
	c := Conversation{}
	err := Conversations.Find(bson.M{"token": token}).One(&c)

	return &c, err
}