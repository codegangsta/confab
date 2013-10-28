package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
)

var DB *mgo.Database

func InitDB() {
	name := os.Getenv("DB_NAME")

	session, err := mgo.Dial(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	println("connecting to db:", name)
	DB = session.DB(name)
	Conversations = DB.C("conversations")
}

var Conversations *mgo.Collection

type Conversation struct {
	Token  string `json:"token"`
	Email1 string `json:"email1"`
	Email2 string `json:"email2"`
	Name1  string `json:"name1"`
	Name2  string `json:"name2"`
}

func CreateConversation(email1 string, name1 string, email2 string, name2 string) (*Conversation, error) {
	token := NewUUID().String()

	c := Conversation{
		Token:  token,
		Email1: email1,
		Email2: email2,
		Name1:  name1,
		Name2:  name2,
	}

	err := Conversations.Insert(&c)

	return &c, err
}

func GetConversation(token string) (*Conversation, error) {
	c := Conversation{}
	err := Conversations.Find(bson.M{"token": token}).One(&c)

	return &c, err
}
